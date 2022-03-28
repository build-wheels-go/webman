package distributed

import (
	"webman/framework"
	"webman/framework/contract"
)

var _ framework.ServiceProvider = (*LocalDistributedProvider)(nil)

type LocalDistributedProvider struct {
}

func (sp *LocalDistributedProvider) Register(container framework.Container) framework.NewInstance {
	return NewLocalDistributedService
}

func (sp *LocalDistributedProvider) Boot(container framework.Container) error {
	return nil
}

func (sp *LocalDistributedProvider) IsDefer() bool {
	return false
}

func (sp *LocalDistributedProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

func (sp *LocalDistributedProvider) Name() string {
	return contract.DistributedKey
}
