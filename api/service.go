// Package api
package api

const (
	// required signature NewPluginService(*errgroup.Group, context.Context, api.IMessenger) api.IPluginService
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
