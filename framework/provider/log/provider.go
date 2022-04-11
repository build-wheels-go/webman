package log

import (
	"webman/framework"
	"webman/framework/contract"
)

var _ framework.ServiceProvider = (*LogProvider)(nil)

type LogProvider struct {
}

func (provider *LogProvider) Register(container framework.Container) framework.NewInstance {
	return nil
}

func (provider *LogProvider) Boot(container framework.Container) error {
	return nil
}

func (provider *LogProvider) IsDefer() bool {
	return false
}

func (provider *LogProvider) Params(container framework.Container) []interface{} {
	return nil
}

func (provider *LogProvider) Name() string {
	return contract.LogKey
}
