// Package core
package core

import (
	"context"
	"log/slog"

	"gomodss/api"

	"golang.org/x/sync/errgroup"
)

type MessengerService struct {
	name     string
	ctx      context.Context
	jobQueue api.IJobQueue
	handlers map[api.Topic][]api.IMessengerHandler
}

func NewMessengerService(eg *errgroup.Group, ctx context.Context) MessengerService {
	return MessengerService{
		name:     "MessengerService",
		ctx:      ctx,
		jobQueue: api.NewJobQueue("messenger", eg),
		handlers: map[api.Topic][]api.IMessengerHandler{},
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
	service.jobQueue.Start(service.ctx)
}

func (service MessengerService) Stop() {
	slog.Info("STOP MessengerService")
}

func (service MessengerService) Publish(topic api.Topic, message api.Message, callback api.IPublishCallback) {
	slog.Info("publish", "topic", topic, "message", message)
	callback.OnSuccess(api.StatusOK, 200)
}

func (service MessengerService) Subscribe(topic api.Topic, listener api.ISubscribeListener) {
}

func (service MessengerService) Unsubscribe(topic api.Topic, listener api.ISubscribeListener) {
}

func (service MessengerService) Register(t api.Topic, handler api.IMessengerHandler) {
	// TODO: use jobqueue here
	slog.Info("AddHandler")
	topicHandlers, ok := service.handlers[t]
	if ok {
		topicHandlers = append(topicHandlers, handler)
		service.handlers[t] = topicHandlers
	} else {
		service.handlers[t] = []api.IMessengerHandler{handler}
	}
	slog.Info("handler added for", "topic", t, "handler", handler)
}

func (service MessengerService) Unregister(t api.Topic, handler api.IMessengerHandler) {
	// TODO: use jobqueue here
	slog.Info("RemoveHandler")
}
