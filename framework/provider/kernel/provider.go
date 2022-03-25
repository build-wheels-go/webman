package kernel

import (
	"webman/framework"
	"webman/framework/contract"
	"webman/framework/gin"
)

var _ framework.ServiceProvider = (*WmKernelProvider)(nil)

type WmKernelProvider struct {
	HttpEngine *gin.Engine
}

func (provider *WmKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewWmKernelService
}

func (provider *WmKernelProvider) Name() string {
	return contract.KernelKey
}

func (provider *WmKernelProvider) IsDefer() bool {
	return false
}

func (provider *WmKernelProvider) Boot(c framework.Container) error {
	if provider.HttpEngine == nil {
		provider.HttpEngine = gin.Default()
	}
	provider.HttpEngine.SetContainer(c)
	return nil
}

func (provider *WmKernelProvider) Params(c framework.Container) []interface{} {
	return []interface{}{provider.HttpEngine}
}
