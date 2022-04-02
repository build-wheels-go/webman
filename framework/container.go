package framework

import (
	"errors"
	"sync"
)

type Container interface {
	// Bind 绑定一个服务提供者，如果凭证已存在则替换
	Bind(sp ServiceProvider) error
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

func NewWmContainer() *WmContainer {
	return &WmContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

func (c *WmContainer) Bind(sp ServiceProvider) error {
	c.lock.Lock()
	key := sp.Name()
	c.providers[key] = sp
	c.lock.Unlock()
	if !sp.IsDefer() {
		params := sp.Params(c)
		instance, err := c.newInstance(sp, params)
		if err != nil {
			return err
		}
		c.instances[key] = instance
	}
	return nil
}

func (c *WmContainer) IsBind(key string) bool {
	return c.findServiceProvider(key) != nil
}

func (c *WmContainer) findServiceProvider(key string) ServiceProvider {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if p, ok := c.providers[key]; ok {
		return p
	}
	return nil
}

func (c *WmContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	if err := sp.Boot(c); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(c)
	}
	method := sp.Register(c)
	ins, err := method(params...)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

func (c *WmContainer) Make(key string) (interface{}, error) {
	return c.make(key, nil, false)
}

func (c *WmContainer) MustMake(key string) interface{} {
	ins, err := c.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return ins
}

func (c *WmContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return c.make(key, params, true)
}

func (c *WmContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	sp := c.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}
	if forceNew {
		return c.newInstance(sp, params)
	}
	if ins, ok := c.instances[key]; ok {
		return ins, nil
	}
	ins, err := c.newInstance(sp, params)
	if err != nil {
		return nil, err
	}
	c.instances[key] = ins

	return ins, nil
}
