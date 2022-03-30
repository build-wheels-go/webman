package main

import (
	"log"
	"webman/app/console"
	"webman/app/http"
	"webman/framework"
	"webman/framework/provider/app"
	"webman/framework/provider/distributed"
	"webman/framework/provider/kernel"
)

func main() {
	//初始化服务容器
	container := framework.NewWmContainer()
	_ = container.Bind(&app.WmAppProvider{})

	_ = container.Bind(&distributed.LocalDistributedProvider{})

	if r, err := http.NewHttpEngine(); err == nil {
		_ = container.Bind(&kernel.WmKernelProvider{HttpEngine: r})
	}

	// 运行命令
	if err := console.RunCommand(container); err != nil {
		log.Fatal("Run Command:", err)
	}
}
