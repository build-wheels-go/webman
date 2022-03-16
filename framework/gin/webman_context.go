// Copyright 2021 build-wheels-go.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package gin

import (
	"context"
)

func (ctx *Context) BaseContext() context.Context {
	return ctx.Request.Context()
}
