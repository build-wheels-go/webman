package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"webman/framework"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), 5*time.Second)
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		Foo(c)
		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "timeout")
		c.SetHasTimeout()

	}
	return nil
}

func Foo(ctx *framework.Context) error {
	obj := map[string]interface{}{
		"errno":  50001,
		"errmsg": "inner error",
		"data":   nil,
	}

	fooInt := ctx.FormInt("foo", 10)
	obj["data"] = fooInt
	return ctx.Json(http.StatusOK, obj)
}
