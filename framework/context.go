package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	//中间件链条
	handlers []ControllerHandler
	//当前中间件的节点
	index int
	//超时标记符
	hasTimeout bool
	//写保护机制
	writerMux *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
		index:          -1,
	}
}

// #region base function

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

// #endregion

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

// #region implement context.Context

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

// #region

// #region query url

func (ctx *Context) QueryInt(k string, def int) int {
	params := ctx.QueryAll()
	if v, ok := params[k]; ok {
		l := len(v)
		if l > 0 {
			intV, err := strconv.Atoi(v[l-1])
			if err != nil {
				return def
			}
			return intV
		}
	}
	return def
}

func (ctx *Context) QueryString(k string, def string) string {
	params := ctx.QueryAll()
	if v, ok := params[k]; ok {
		l := len(v)
		if l > 0 {
			return v[l-1]
		}
	}
	return def
}

func (ctx *Context) QueryArray(k string, def []string) []string {
	params := ctx.QueryAll()
	if v, ok := params[k]; ok {
		return v
	}
	return def
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.URL.Query()
	}
	return map[string][]string{}
}

// #endregion

// #region form post

func (ctx *Context) FormInt(k string, def int) int {
	params := ctx.FormAll()
	if v, ok := params[k]; ok {
		l := len(v)
		if l > 0 {
			intV, err := strconv.Atoi(v[l-1])
			if err != nil {
				return def
			}
			return intV
		}
	}
	return def
}

func (ctx *Context) FormString(k string, def string) string {
	params := ctx.FormAll()
	if v, ok := params[k]; ok {
		l := len(v)
		if l > 0 {
			return v[l-1]
		}
	}
	return def
}

func (ctx *Context) FormArray(k string, def []string) []string {
	params := ctx.FormAll()
	if v, ok := params[k]; ok {
		return v
	}
	return def
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil {
		return ctx.request.PostForm
	}
	return map[string][]string{}
}

// #endregion

// #region application/json post

func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request == nil {
		return errors.New("ctx.request is empty")
	}
	body, err := ioutil.ReadAll(ctx.request.Body)
	if err != nil {
		return err
	}
	ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	if err := json.Unmarshal(body, obj); err != nil {
		return err
	}
	return nil
}

// #endregion

// #region response

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.hasTimeout {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	if _, err = ctx.responseWriter.Write(byt); err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	return nil
}

func (ctx *Context) Html(status int, obj interface{}, template string) error {
	return nil
}

func (ctx *Context) Text() error {
	return nil
}

// #endregion
