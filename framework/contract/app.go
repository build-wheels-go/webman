package contract

const Key = "wm:app"

type App interface {
	// Version 当前版本
	Version() string
	// BaseFolder 定义基础路径
	BaseFolder() string
	// ConfigFolder 定义配置路径
	ConfigFolder() string
	// ProviderFolder 定义业务服务提供者路径
	ProviderFolder() string
	// MiddlewareFolder 定义业务中间件路径
	MiddlewareFolder() string
	// LogFolder 定义日志路径
	LogFolder() string
	// CommandFolder 定义业务命令路径
	CommandFolder() string
	// RuntimeFolder 定义运行中间态路径
	RuntimeFolder() string
	// TestFolder 定义测试文件路径
	TestFolder() string
}
