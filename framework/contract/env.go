package contract

const EnvKey = "wm:env"

const (
	// 生产环境
	EnvProduction = "production"
	// 测试环境
	EnvTesting = "testing"
	// 开发环境
	EnvDevelopment = "development"
)

type Env interface {
	// 获取当前的环境
	AppEnv() string
	// 判断环境变量是否被设置
	IsExist(key string) bool
	// 获取单个环境变量
	Get(key string) string
	// 获取全部环境变量
	All() map[string]string
}