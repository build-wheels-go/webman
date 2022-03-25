package cobra

import "webman/framework"

// SetContainer 设置服务容器
func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

// GetContainer 获取服务容器
func (c *Command) GetContainer() framework.Container {
	return c.Root().container
}
