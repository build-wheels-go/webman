package main

import (
	"webman/framework/gin"
	"webman/provider/demo"
)

func UserLoginController(c *gin.Context) {
	demoService := c.MustMake(demo.Key).(demo.Service)
	foo := demoService.GetFoo()
	c.ISetOkStatus().IJson(foo)
}
