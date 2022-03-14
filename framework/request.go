package framework

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"mime/multipart"

	"github.com/spf13/cast"
)

const defaultMultipartMemory = 32 << 20 //32MB

type IRequest interface {
	// 请求地址url中的参数
	// 如: foo.com?a=1&b=bar&c[]=bar
	QueryInt(key string, def int) (int, bool)
	QueryInt64(key string, def int64) (int64, bool)
	QueryFloat64(key string, def float64) (float64, bool)
	QueryFloat32(key string, def float32) (float32, bool)
	QueryBool(key string, def bool) (bool, bool)
	QueryString(key string, def string) (string, bool)
	QueryStringSlice(key string, def []string) ([]string, bool)
	Query(key string) interface{}

	// 路由匹配中带点参数
	// 如 /foo/:id
	ParamInt(key string, def int) (int, bool)
	ParamInt64(key string, def int64) (int64, bool)
	ParamFloat64(key string, def float64) (float64, bool)
	ParamFloat32(key string, def float32) (float32, bool)
	ParamBool(key string, def bool) (bool, bool)
	ParamString(key string, def string) (string, bool)
	Param(key string) interface{}

	// form表单中带的参数
	FormInt(key string, def int) (int, bool)
	FormInt64(key string, def int64) (int64, bool)
	FormFloat64(key string, def float64) (float64, bool)
	FormFloat32(key string, def float32) (float32, bool)
	FormBool(key string, def bool) (bool, bool)
	FormString(key string, def string) (string, bool)
	FormStringSlice(key string, def []string) ([]string, bool)
	FormFile(key string) (*multipart.FileHeader, error)
	Form(key string) interface{}

	// json body
	BindJson(obj interface{}) error

	// xml body
	BindXml(obj interface{}) error

	// 其他格式
	GetRowData() ([]byte, error)

	// 基础信息
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// header
	Headers() map[string][]string
	Header(key string) (string, bool)

	// cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

// QueryAll 获取请求地址中的所有参数
func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

func (ctx *Context) QueryInt(key string, def int) (int, bool) {
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

func (ctx *Context) QueryInt64(key string, def int64) (int64, bool) {
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

func (ctx *Context) QueryFloat64(key string, def float64) (float64, bool) {
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

func (ctx *Context) QueryFloat32(key string, def float32) (float32, bool) {
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

func (ctx *Context) QueryBool(key string, def bool) (bool, bool) {
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

func (ctx *Context) QueryString(key string, def string) (string, bool) {
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

func (ctx *Context) QueryStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

func (ctx *Context) Query(key string) interface{} {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals[0]
	}
	return nil
}

func (ctx *Context) ParamInt(key string, def int) (int, bool) {
	if val := ctx.Param(key); val != nil {
		val, err := cast.ToIntE(val)
		if err != nil {
			return def, false
		}
		return val, true
	}
	return def, false
}

func (ctx *Context) ParamInt64(key string, def int64) (int64, bool) {
	if val := ctx.Param(key); val != nil {
		val, err := cast.ToInt64E(val)
		if err != nil {
			return def, false
		}
		return val, true
	}
	return def, false
}

func (ctx *Context) ParamFloat64(key string, def float64) (float64, bool) {
	if val := ctx.Param(key); val != nil {
		val, err := cast.ToFloat64E(val)
		if err != nil {
			return def, false
		}
		return val, true
	}
	return def, false
}

func (ctx *Context) ParamFloat32(key string, def float32) (float32, bool) {
	if val := ctx.Param(key); val != nil {
		val, err := cast.ToFloat32E(val)
		if err != nil {
			return def, false
		}
		return val, true
	}
	return def, false
}

func (ctx *Context) ParamBool(key string, def bool) (bool, bool) {
	if val := ctx.Param(key); val != nil {
		val, err := cast.ToBoolE(val)
		if err != nil {
			return def, false
		}
		return val, true
	}
	return def, false
}

func (ctx *Context) ParamString(key string, def string) (string, bool) {
	if val := ctx.Param(key); val != nil {
		val, err := cast.ToStringE(val)
		if err != nil {
			return def, false
		}
		return val, true
	}
	return def, false
}

func (ctx *Context) Param(key string) interface{} {
	if ctx.params != nil {
		if val, ok := ctx.params[key]; ok {
			return val
		}
	}
	return nil
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) (int, bool) {
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

func (ctx *Context) FormInt64(key string, def int64) (int64, bool) {
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

func (ctx *Context) FormFloat64(key string, def float64) (float64, bool) {
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

func (ctx *Context) FormFloat32(key string, def float32) (float32, bool) {
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

func (ctx *Context) FormBool(key string, def bool) (bool, bool) {
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

func (ctx *Context) FormString(key string, def string) (string, bool) {
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

func (ctx *Context) FormStringSlice(key string, def []string) ([]string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

func (ctx *Context) FormFile(key string) (*multipart.FileHeader, error) {
	if ctx.request.MultipartForm != nil {
		if err := ctx.request.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := ctx.request.FormFile(key)
	if err != nil {
		return nil, err
	}
	f.Close()

	return fh, nil
}

func (ctx *Context) Form(key string) interface{} {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals[0]
	}
	return nil
}

// json body
func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request == nil {
		return errors.New("context.request empty")
	}
	body, err := ioutil.ReadAll(ctx.request.Body)
	if err != nil {
		return err
	}
	// 重新填充request.Body，为后续的逻辑二次读取做准备
	ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	if err = json.Unmarshal(body, obj); err != nil {
		return err
	}
	return nil
}

// xml body
func (ctx *Context) BindXml(obj interface{}) error {
	if ctx.request == nil {
		return errors.New("context.request empty")
	}
	body, err := ioutil.ReadAll(ctx.request.Body)
	if err != nil {
		return err
	}
	// 重新填充request.Body，为后续的逻辑二次读取做准备
	ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	if err = xml.Unmarshal(body, obj); err != nil {
		return err
	}
	return nil
}

// 其他格式
func (ctx *Context) GetRowData() ([]byte, error) {
	if ctx.request == nil {
		return nil, errors.New("context.request empty")
	}
	body, err := ioutil.ReadAll(ctx.request.Body)
	if err != nil {
		return nil, err
	}
	// 重新填充request.Body，为后续的逻辑二次读取做准备
	ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return body, nil
}

func (ctx *Context) Uri() string {
	return ctx.request.RequestURI
}

func (ctx *Context) Method() string {
	return ctx.request.Method
}

func (ctx *Context) Host() string {
	return ctx.request.URL.Host
}

func (ctx *Context) ClientIp() string {
	ipAddr := ctx.request.Header.Get("X-Real-Ip")
	if ipAddr == "" {
		ipAddr = ctx.request.Header.Get("X-Forwarded-For")
	}
	if ipAddr == "" {
		ipAddr = ctx.request.RemoteAddr
	}

	return ipAddr
}

// header
func (ctx *Context) Headers() map[string][]string {
	return ctx.request.Header
}

func (ctx *Context) Header(key string) (string, bool) {
	vals := ctx.request.Header.Values(key)
	if vals == nil || len(vals) == 0 {
		return "", false
	}
	return vals[0], true
}

// cookie
func (ctx *Context) Cookies() map[string]string {
	cookies := ctx.request.Cookies()
	ret := map[string]string{}
	for _, c := range cookies {
		ret[c.Name] = c.Value
	}
	return ret
}

func (ctx *Context) Cookie(key string) (string, bool) {
	cookies := ctx.Cookies()
	if val, ok := cookies[key]; ok {
		return val, true
	}
	return "", false
}
