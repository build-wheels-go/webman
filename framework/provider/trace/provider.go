package trace

import (
	"webman/framework"
	"webman/framework/contract"
)

var _ framework.ServiceProvider = (*WmTraceProvider)(nil)

type WmTraceProvider struct {
	container framework.Container
}

func (provider *WmTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewWmTraceService
}

func (provider *WmTraceProvider) Boot(c framework.Container) error {
	provider.container = c
	return nil
}

func (provider *WmTraceProvider) IsDefer() bool {
	return false
}

func (provider *WmTraceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (provider *WmTraceProvider) Name() string {
	return contract.TraceKey
}
