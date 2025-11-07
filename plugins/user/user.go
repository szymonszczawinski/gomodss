// Package user
package main

import (
	"log/slog"

	"gomodss/api"
)

type UserService struct {
	messenger api.IMessenger
	name      string
}

func NewPluginService(messenger api.IMessenger) api.IPluginService {
	slog.Info("INIT UserService")
	return UserService{
		messenger: messenger,
		name:      "UserService",
	}
}

func (service UserService) String() string {
	return service.name
}

func (service UserService) Create() {
	slog.Info("CREATE UserService")
}

func (service UserService) Start() {
	slog.Info("START UserService")
	callback := loginCallback{}
	service.messenger.Publish(api.UserLogin, "login now", callback)
}

func (service UserService) Stop() {
	slog.Info("STOP UserService")
}

func (service UserService) NewModService(messenger api.IMessenger) {
	slog.Info("INIT UserService")
}

type loginCallback struct{}

func (c loginCallback) OnSuccess(status api.PublishStatus, code int) {
	slog.Info("on sucess login", "status", status, "code", code)
}

func (c loginCallback) OnError(status api.PublishStatus, code int, errorMessage string) {}

func main() {}
