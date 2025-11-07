// Package notification
package main

import (
	"log/slog"

	"gomodss/api"
)

type NotificationService struct {
	name      string
	messenger api.IMessenger
}

func NewPluginService(messenger api.IMessenger) api.IPluginService {
	slog.Info("INIT NotificationService")
	return NotificationService{
		messenger: messenger,
		name:      "NotificationService",
	}
}

func (service NotificationService) String() string {
	return service.name
}

func (service NotificationService) Create() {
	slog.Info("CREATE NotificationService")
}

func (service NotificationService) Start() {
	slog.Info("START NotificationService")
}

func (service NotificationService) Stop() {
	slog.Info("STOP NotificationService")
}

func main() {}
