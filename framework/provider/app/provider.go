package app

import (
	"webman/framework"
	"webman/framework/contract"
)

var _ framework.ServiceProvider = (*WmAppProvider)(nil)

type WmAppProvider struct {
	BaseFolder string
}

func (provider *WmAppProvider) Register(c framework.Container) framework.NewInstance {
	return NewWmApp
}

func (provider *WmAppProvider) Name() string {
	return contract.Key
}

func (provider *WmAppProvider) IsDefer() bool {
	return false
}

func (provider *WmAppProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c, provider.BaseFolder}
}

func (provider *WmAppProvider) Boot(c framework.Container) error {
	return nil
}
