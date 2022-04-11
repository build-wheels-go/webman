package id

import (
	"webman/framework"
	"webman/framework/contract"
)

var _ framework.ServiceProvider = (*WmIDProvider)(nil)

type WmIDProvider struct {
}

func (id *WmIDProvider) Register(c framework.Container) framework.NewInstance {
	return NewIDService
}

func (id *WmIDProvider) Boot(c framework.Container) error {
	return nil
}

func (id *WmIDProvider) IsDefer() bool {
	return false
}

func (id *WmIDProvider) Params(c framework.Container) []interface{} {
	return []interface{}{}
}

func (id *WmIDProvider) Name() string {
	return contract.IDKey
}
