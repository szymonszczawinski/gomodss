package api

const (
	StatusOK    PublishStatus = "OK"
	StatusERROR PublishStatus = "ERROR"

	// User Topics

	UserLogin    Topic = "UserLogin"
	UserLoginAck Topic = "UserLoginAck"

	// Notification Topics

	NotificationError    Topic = "NotificationError"
	NotificationErrorAck Topic = "NotificationErrorAck"
)

type (
	PublishStatus string
	Topic         string
	Message       any
)

// IMessenger used as input interface for messenger
type IMessenger interface {
	// publish message on topic with callback; select appropriate IMessengerHandler based on topic
	// and forward message via handler::OnPublish
	Publish(topic Topic, message Message, callback IPublishCallback)
	// select appropriate IMessengerHandler based on topic and subscribe via handler::OnSubscribe
	Subscribe(topic Topic, listener ISubscribeListener)
	// select appropriate IMessengerHandler based on topic and unsubscribe via handler::OnUnsubscribe
	Unsubscribe(topic Topic, listener ISubscribeListener)
}

// IMessengerRegistry used to register and unregister IMessengerHandler for given topic
type IMessengerRegistry interface {
	Register(t Topic, handler IMessengerHandler)
	Unregister(t Topic, handler IMessengerHandler)
}

type IPublishCallback interface {
	OnSuccess(status PublishStatus, code int)
	OnError(status PublishStatus, code int, errorMessage string)
}

type ISubscribeListener interface {
	OnMessage(topic Topic, message Message)
}

// IMessengerHandler defines api for messenger out enpoints;
// Particular component need to implement this to be able to handle pubslis/subscribe requests;
// Such component need to register itself with IMessengerRegistry::Register
type IMessengerHandler interface {
	OnPublish(t Topic, m Message, callback IPublishCallback)
	OnSubscribe(t Topic, listener ISubscribeListener)
	OnUnsubscribe(t Topic, listener ISubscribeListener)
}
