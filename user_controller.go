package main

import (
	"webman/framework/gin"
)

func UserLoginController(c *gin.Context) {
	c.ISetOkStatus().IJson(map[string]string{"data": "ok, UserLoginController"})
}
