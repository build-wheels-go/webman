package demo

import (
	"webman/app/provider/demo"
	"webman/framework/contract"
	"webman/framework/gin"
)

type DemoApi struct {
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	r.Bind(&demo.DemoServiceProvider{})
	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/config", api.Config)
	return nil
}

func NewDemoApi() *DemoApi {
	return &DemoApi{}
}

func (api *DemoApi) Demo(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"data": "ok"})
}

func (api *DemoApi) Config(ctx *gin.Context) {
	configService := ctx.MustMake(contract.ConfigKey).(contract.Config)
	data := map[string]string{
		"hostname": configService.GetString("db.mysql.hostname"),
		"username": configService.GetString("db.mysql.username"),
		"password": configService.GetString("db.mysql.password"),
	}
	ctx.JSON(200, data)
}
