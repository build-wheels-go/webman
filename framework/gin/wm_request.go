// Copyright 2021 build-wheels-go.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package gin

import (
	"github.com/spf13/cast"
)

type IRequest interface {
	// 请求地址url中的参数
	// 如: foo.com?a=1&b=bar&c[]=bar
	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)

	// 路由匹配中带点参数
	// 如 /foo/:id
	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)

	// form表单中带的参数
	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultFormStringSlice(key string, def []string) ([]string, bool)
	DefaultForm(key string) interface{}
}

// QueryAll 获取请求地址中的所有参数
func (ctx *Context) QueryAll() map[string][]string {
	ctx.initQueryCache()
	return ctx.queryCache
}

func (ctx *Context) DefaultQueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			val, err := cast.ToIntE(vals[0])
			if err != nil {
				return def, false
			}
			return val, true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryInt64(key string, def int64) (int64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			val, err := cast.ToInt64E(vals[0])
			if err != nil {
				return def, false
			}
			return val, true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			val, err := cast.ToFloat64E(vals[0])
			if err != nil {
				return def, false
			}
			return val, true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			val, err := cast.ToFloat32E(vals[0])
			if err != nil {
				return def, false
			}
			return val, true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			val, err := cast.ToBoolE(vals[0])
			if err != nil {
				return def, false
			}
			return val, true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryString(key string, def string) (string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			val, err := cast.ToStringE(vals[0])
			if err != nil {
				return def, false
			}
			return val, true
		}
	}
	return def, false
}

func (ctx *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

func (ctx *Context) DefaultParamInt(key string, def int) (int, bool) {
	if val := ctx.WmParam(key); val != nil {
		val, err := cast.ToIntE(val)
		if err != nil {
			return def, false
		}
		return val, true
	}
	return def, false
}

func (ctx *Context) DefaultParamInt64(key string, def int64) (int64, bool) {
	if val := ctx.WmParam(key); val != nil {
		val, err := cast.ToInt64E(val)
		if err != nil {
			return def, false
		}
		return val, true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat64(key string, def float64) (float64, bool) {
	if val := ctx.WmParam(key); val != nil {
		val, err := cast.ToFloat64E(val)
		if err != nil {
			return def, false
		}
		return val, true
	}
	return def, false
}

func (ctx *Context) DefaultParamFloat32(key string, def float32) (float32, bool) {
	if val := ctx.WmParam(key); val != nil {
		val, err := cast.ToFloat32E(val)
		if err != nil {
			return def, false
		}
		return val, true
	}
	return def, false
}

func (ctx *Context) DefaultParamBool(key string, def bool) (bool, bool) {
	if val := ctx.WmParam(key); val != nil {
		val, err := cast.ToBoolE(val)
		if err != nil {
			return def, false
		}
		return val, true
	}
	return def, false
}

func (ctx *Context) DefaultParamString(key string, def string) (string, bool) {
	if val := ctx.WmParam(key); val != nil {
		val, err := cast.ToStringE(val)
		if err != nil {
			return def, false
		}
		return val, true
	}
	return def, false
}

func (ctx *Context) WmParam(key string) interface{} {
	if val, ok := ctx.Params.Get(key); ok {
		return val
	}

	return nil
}

func (ctx *Context) FormAll() map[string][]string {
	ctx.initFormCache()
	return ctx.formCache
}

func (ctx *Context) DefaultFormInt(key string, def int) (int, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			val, err := cast.ToIntE(vals[0])
			if err != nil {
				return def, false
			}
			return val, true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormInt64(key string, def int64) (int64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			val, err := cast.ToInt64E(vals[0])
			if err != nil {
				return def, false
			}
			return val, true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormFloat64(key string, def float64) (float64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			val, err := cast.ToFloat64E(vals[0])
			if err != nil {
				return def, false
			}
			return val, true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormFloat32(key string, def float32) (float32, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			val, err := cast.ToFloat32E(vals[0])
			if err != nil {
				return def, false
			}
			return val, true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormBool(key string, def bool) (bool, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			val, err := cast.ToBoolE(vals[0])
			if err != nil {
				return def, false
			}
			return val, true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormString(key string, def string) (string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			val, err := cast.ToStringE(vals[0])
			if err != nil {
				return def, false
			}
			return val, true
		}
	}
	return def, false
}

func (ctx *Context) DefaultFormStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

func (ctx *Context) DefaultForm(key string) interface{} {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals[0]
	}
	return nil
}
