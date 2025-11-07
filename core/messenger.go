// Package core
package core

import (
	"log/slog"

	"gomodss/api"
)

type MessengerService struct {
	name string
}

func NewMessengerService() MessengerService {
	return MessengerService{
		name: "MessengerService",
	}
}

func (service MessengerService) String() string {
	return service.name
}

func (service MessengerService) Create() {
	slog.Info("CREATE MessengerService")
}

func (service MessengerService) Start() {
	slog.Info("START MessengerService")
}

func (service MessengerService) Stop() {
	slog.Info("STOP MessengerService")
}

func (service MessengerService) Publish(topic api.Topic, message any, callback api.IPublishCallback) {
	slog.Info("publish", "topic", topic, "message", message)
	callback.OnSuccess(api.StatusOK, 200)
}

func (service MessengerService) Subscribe(topic api.Topic, listener api.ISubscribeListener) {
}
