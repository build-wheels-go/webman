package demo

import (
	"fmt"
	"webman/framework"
)

var _ framework.ServiceProvider = (*DemoServiceProvider)(nil)

type DemoServiceProvider struct {
}

func (demo *DemoServiceProvider) Name() string {
	return Key
}

func (demo *DemoServiceProvider) Register(c framework.Container) framework.NewInstance {
	return NewDemoService
}

func (demo *DemoServiceProvider) IsDefer() bool {
	return true
}

func (demo *DemoServiceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (demo *DemoServiceProvider) Boot(c framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}
