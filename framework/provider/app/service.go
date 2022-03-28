package app

import (
	"errors"
	"flag"
	"path/filepath"
	"webman/framework"
	"webman/framework/contract"
	"webman/framework/util"

	"github.com/google/uuid"
)

var _ contract.App = (*WmApp)(nil)

type WmApp struct {
	container  framework.Container //服务容器
	baseFolder string              //基础路径
	appId      string              //表示当前app的唯一标识，可用于分布式锁
}

func NewWmApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	if baseFolder == "" {
		flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数，默认为当前路径")
		flag.Parse()
	}

	appId := uuid.New().String()
	return &WmApp{container: container, baseFolder: baseFolder, appId: appId}, nil
}

func (app *WmApp) AppID() string {
	return app.appId
}

func (app *WmApp) Version() string {
	return "0.0.1"
}

func (app *WmApp) BaseFolder() string {
	if app.baseFolder != "" {
		return app.baseFolder
	}
	// 参数也没有使用当前路径
	return util.GetExecDirectory()
}

func (app *WmApp) ConfigFolder() string {
	return filepath.Join(app.BaseFolder(), "config")
}

func (app *WmApp) HttpFolder() string {
	return filepath.Join(app.BaseFolder(), "http")
}

func (app *WmApp) MiddlewareFolder() string {
	return filepath.Join(app.HttpFolder(), "middleware")
}

func (app *WmApp) StorageFolder() string {
	return filepath.Join(app.BaseFolder(), "storage")
}

func (app *WmApp) LogFolder() string {
	return filepath.Join(app.StorageFolder(), "log")
}

func (app *WmApp) RuntimeFolder() string {
	return filepath.Join(app.StorageFolder(), "runtime")
}

func (app *WmApp) ConsoleFolder() string {
	return filepath.Join(app.BaseFolder(), "console")
}

func (app *WmApp) CommandFolder() string {
	return filepath.Join(app.ConsoleFolder(), "command")
}

func (app *WmApp) ProviderFolder() string {
	return filepath.Join(app.BaseFolder(), "provider")
}

func (app *WmApp) TestFolder() string {
	return filepath.Join(app.BaseFolder(), "test")
}
