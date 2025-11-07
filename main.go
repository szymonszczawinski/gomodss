// Package gomod
package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"os"
	"plugin"

	"gomodss/api"
	"gomodss/core"
)

func main() {
	services := []api.IService{}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	slog.Info("starting GOMODSS")

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
	slog.Info("----------------------------")

	slog.Info("stoping GOMODSS")
	for _, s := range services {
		s.Stop()
	}
	slog.Info("exiting GOMODSS")
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
		plg, err := plugin.Open("plugins/" + p)
		if err != nil {
			// NOTE: in real cases, deal with this error
			slog.Error("could not open file", "file", p, "err", err)
			continue
		}
		pluginList = append(pluginList, plg)
	}
	return pluginList
}
