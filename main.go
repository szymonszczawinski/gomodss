// Package gomod
package main

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"plugin"
	"syscall"

	"gomodss/api"
	"gomodss/core"

	"golang.org/x/sync/errgroup"
)

var services []api.IService

func main() {
	cancelContext, cancel := context.WithCancel(context.Background())
	signalChannel := registerShutdownHook(cancel)
	mainGroup, groupContext := errgroup.WithContext(cancelContext)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	slog.Info("starting GOMODSS")

	mainQueue := api.NewJobQueue("gomodss", mainGroup)
	mainQueue.Start(groupContext)

	services = []api.IService{}
	startServices(mainGroup, groupContext)
	slog.Info("----------------------------")

	if err := mainGroup.Wait(); err == nil {
		slog.Info("stopping GOMODSS")
	}

	defer close(signalChannel)
	slog.Info("stoping services")
	for _, s := range services {
		s.Stop()
	}
	slog.Info("exiting GOMODSS")
}

func registerShutdownHook(cancel context.CancelFunc) chan os.Signal {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		// wait until receiving the signal
		<-stop
		slog.Info("Shutdown SIGNAL received -> cancel context")
		cancel()
	}()

	return stop
}

func startServices(eg *errgroup.Group, ctx context.Context) {
	messenger := core.NewMessengerService()
	messenger.Create()
	messenger.Start()
	services = append(services, messenger)

	pluginPathList := loadConfig()
	slog.Info("loaded plugins paths", "paths", pluginPathList)
	plugins := loadPlugins(pluginPathList)
	slog.Info("plugins loaded", "plugins", plugins)
	pluginServices := runServices(plugins, messenger)
	services = append(services, pluginServices...)
	slog.Info("all services running", "services", services)
}

func runServices(plugins []*plugin.Plugin, messenger api.IMessenger) []api.IService {
	modServices := []api.IService{}
	for _, plugin := range plugins {
		create, err := plugin.Lookup(api.NewPluginServiceFunction)
		if err != nil {
			slog.Error("could not find 'NewModService' function for plugin", "plugin", plugin, "err", err)
			continue
		}
		newFunction, ok := create.(func(api.IMessenger) api.IPluginService)
		if !ok {
			slog.Error("not a proper New function", "func", create)
		}
		modService := newFunction(messenger)
		modService.Create()
		slog.Info("service created", "service", modService)
		modService.Start()
		slog.Info("service started", "service", modService)
		modServices = append(modServices, modService)
	}
	return modServices
}

func loadConfig() []string {
	var pluginPathList []string
	f, err := os.ReadFile("plugins.json")
	if err != nil {
		// NOTE: in real cases, deal with this error
		panic(err)
	}
	json.Unmarshal(f, &pluginPathList)
	return pluginPathList
}

func loadPlugins(pluginPathList []string) []*plugin.Plugin {
	var pluginList []*plugin.Plugin
	// Allocate a list for storing all our plugins
	pluginList = make([]*plugin.Plugin, 0, len(pluginPathList))
	for _, p := range pluginPathList {
		// We use plugin.Open to load the plugin by path
		plg, err := plugin.Open(p)
		if err != nil {
			// NOTE: in real cases, deal with this error
			slog.Error("could not open file", "file", p, "err", err)
			continue
		}
		pluginList = append(pluginList, plg)
	}
	return pluginList
}
