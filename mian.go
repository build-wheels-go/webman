package main

import (
	"net/http"
	"webman/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)

	serve := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	serve.ListenAndServe()
}
