package main

import (
	"webman/framework/contract"
	"webman/framework/gin"
)

func UserLoginController(c *gin.Context) {
	app := c.MustMake(contract.Key).(contract.App)
	baseFolder := app.BaseFolder()

	c.ISetOkStatus().IJson(baseFolder)
}
