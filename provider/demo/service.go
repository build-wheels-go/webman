package demo

import "webman/framework"

type DemoService struct {
	//接口实现
	Service

	c framework.Container
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	c := params[0].(framework.Container)
	return &DemoService{
		c: c,
	}, nil
}

func (s *DemoService) GetFoo() Foo {
	return Foo{
		Name: "I am Foo",
	}
}
