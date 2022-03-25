package demo

import (
	"webman/app/provider/demo"
	"webman/framework/gin"
)

type DemoApi struct {
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	r.Bind(&demo.DemoServiceProvider{})
	r.GET("/demo/demo", api.Demo)
	return nil
}

func NewDemoApi() *DemoApi {
	return &DemoApi{}
}

func (api *DemoApi) Demo(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"data": "ok"})
}
