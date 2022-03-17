package framework

import "sync"

type Container interface {
	// Bind 绑定一个服务提供者，如果凭证已存在则替换
	Bind(provider ServiceProvider) error
	// IsBind 根据凭证判断是否已经绑定过服务提供者
	IsBind(key string) bool
	// Make 根据凭证获取一个服务
	Make(key string) (interface{}, error)
	// MustMake 根据凭证获取一个服务。该服务必须存在，否则会panic
	MustMake(key string) interface{}
	// MakeNew 根据凭证和参数获取一个服务，不是单例
	MakeNew(key string, params []interface{}) (interface{}, error)
}

type WmContainer struct {
	Container
	// 存储注册的服务提供者
	providers map[string]ServiceProvider
	// 存储具体实例
	instances map[string]interface{}
	// 读写锁，锁住对容器变更操作
	lock sync.RWMutex
}

func (c *WmContainer) Bind(provider ServiceProvider) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	key := provider.Name()
	c.providers[key] = provider

	if provider.IsDefer() == false {
		if err := provider.Boot(c); err != nil {
			return err
		}
		params := provider.Params(c)
		method := provider.Register(c)
		instance, err := method(params...)
		if err != nil {
			return err
		}
		c.instances[key] = instance
	}
	return nil
}

func (c *WmContainer) Make(key string) (interface{}, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return nil, nil
}
