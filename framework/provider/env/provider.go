package env

import (
	"webman/framework"
	"webman/framework/contract"
)

var _ framework.ServiceProvider = (*WmEnvProvider)(nil)

type WmEnvProvider struct {
	Folder string
}

func (provider *WmEnvProvider) Register(container framework.Container) framework.NewInstance {
	return NewWmEnvService
}

func (provider *WmEnvProvider) Boot(container framework.Container) error {
	app := container.MustMake(contract.AppKey).(contract.App)
	provider.Folder = app.BaseFolder()
	return nil
}

func (provider *WmEnvProvider) IsDefer() bool {
	return false
}

func (provider *WmEnvProvider) Params(container framework.Container) []interface{} {
	return []interface{}{provider.Folder}
}

func (provider *WmEnvProvider) Name() string {
	return contract.EnvKey
}
