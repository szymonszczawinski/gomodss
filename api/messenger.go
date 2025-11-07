package api

const (
	StatusOK    PublishStatus = "OK"
	StatusERROR PublishStatus = "ERROR"

	// User Topics

	UserLogin    Topic = "UserLogin"
	UserLoginAck Topic = "UserLoginAck"

	// Notification Topics

	ErrorNotification    Topic = "ErrorNotification"
	ErrorNotificationAck Topic = "ErrorNotificationAck"
)

type (
	PublishStatus string
	Topic         string
)

type IMessenger interface {
	Publish(topic Topic, message any, callback IPublishCallback)
	Subscribe(topic Topic, listener ISubscribeListener)
}

type IPublishCallback interface {
	OnSuccess(status PublishStatus, code int)
	OnError(status PublishStatus, code int, errorMessage string)
}

type ISubscribeListener interface {
	OnMessage(topic Topic, message any)
}
