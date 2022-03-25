package http

import "webman/framework/gin"

func NewHttpEngine() (*gin.Engine, error) {
	// Mode设置为release
	gin.SetMode(gin.ReleaseMode)
	// 启动默认的web引擎
	r := gin.Default()
	// 绑定路由
	Routes(r)
	return r, nil
}
