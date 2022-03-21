package demo

const Key = "wm:demo"

type Service interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}
