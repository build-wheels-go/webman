package framework

type NewInstance func(...interface{}) (interface{}, error)

type ServiceProvider interface {
	// Register 在服务容器中注册一个服务实例化方法
	Register(Container) NewInstance
	// Boot 在调用实例化服务时的前置操作
	Boot(Container) error
	// IsDefer 是否在注册时就实例化服务
	// false 注册时就实例化服务。true 第一次调用实例化服务
	IsDefer() bool
	// Params 定义传递给NewInstance的参数
	Params(Container) []interface{}
	// Name 服务提供者的凭证
	Name() string
}
