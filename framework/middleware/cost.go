// Copyright 2021 build-wheels-go.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package middleware

import (
	"fmt"
	"time"

	"webman/framework/gin"
)

func Cost() gin.HandlerFunc {
	return func(c *gin.Context) {
		//开始时间
		start := time.Now()
		c.Next()

		//结束时间
		end := time.Now()
		cost := end.Sub(start)
		fmt.Printf("api uri: %v,cost: %v", c.Request.RequestURI, cost.Seconds())
	}
}
