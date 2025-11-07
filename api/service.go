// Package api
package api

const (
	NewPluginServiceFunction = "NewPluginService"
)

type IService interface {
	Create()
	Start()
	Stop()
}
type IPluginService interface {
	IService
}
