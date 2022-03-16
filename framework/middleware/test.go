// Copyright 2021 build-wheels-go.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package middleware

import (
	"fmt"

	"webman/framework/gin"
)

func Test1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware test1 pre")
		c.Next()
		fmt.Println("middleware test1 post")
	}
}

func Test2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware test2 pre")
		c.Next()
		fmt.Println("middleware test2 post")
	}
}
