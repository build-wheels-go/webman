package main

import (
	"time"
	"webman/framework"
)

func UserLoginController(c *framework.Context) error {
	time.Sleep(10 * time.Second)
	c.SetOkStatus().Json(map[string]string{"data": "ok, UserLoginController"})
	return nil
}
