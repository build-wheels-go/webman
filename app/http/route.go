package http

import (
	"webman/app/http/module/demo"
	"webman/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")
	if err := demo.Register(r); err != nil {
		panic(err)
	}
}
