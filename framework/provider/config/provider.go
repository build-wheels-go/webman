package config

import (
	"path/filepath"
	"webman/framework"
	"webman/framework/contract"
)

var _ framework.ServiceProvider = (*WmConfigProvider)(nil)

type WmConfigProvider struct{}

func (provider *WmConfigProvider) Register(container framework.Container) framework.NewInstance {
	return NewWmConfig
}

func (provider *WmConfigProvider) Boot(container framework.Container) error {
	return nil
}

func (provider *WmConfigProvider) IsDefer() bool {
	return false
}

func (provider *WmConfigProvider) Params(container framework.Container) []interface{} {
	appService := container.MustMake(contract.AppKey).(contract.App)
	envService := container.MustMake(contract.EnvKey).(contract.Env)
	env := envService.AppEnv()
	//配置文件夹地址
	configFolder := filepath.Join(appService.ConfigFolder(), env)

	envs := envService.All()
	return []interface{}{container, configFolder, envs}
}

func (provider *WmConfigProvider) Name() string {
	return contract.ConfigKey
}
