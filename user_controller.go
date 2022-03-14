package main

import "webman/framework"

func UserLoginController(c *framework.Context) error {
	c.SetOkStatus().Json(map[string]string{"data": "ok, UserLoginController"})
	return nil
}
