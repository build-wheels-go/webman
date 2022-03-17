// Copyright 2021 build-wheels-go.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package gin

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

type IResponse interface {
	// json 输出
	IJson(obj interface{}) IResponse

	// jsonp 输出
	IJsonp(obj interface{}) IResponse

	// xml 输出
	IXml(obj interface{}) IResponse

	// html 输出
	IHtml(template string, obj interface{}) IResponse

	// 文本输出
	IText(format string, values ...interface{}) IResponse

	// 重定向
	IRedirect(path string) IResponse

	// 设置Header
	ISetHeader(key, val string) IResponse

	// 设置Cookie
	ISetCookie(key, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	ISetStatus(code int) IResponse

	// 设置200状态码
	ISetOkStatus() IResponse
}

func (ctx *Context) IJson(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.ISetHeader("Content-Type", "application/json")
	_, _ = ctx.Writer.Write(byt)
	return ctx
}

func (ctx *Context) IJsonp(obj interface{}) IResponse {
	callbackFunc, _ := ctx.DefaultQueryString("callback", "callback_function")
	ctx.ISetHeader("Content-Type", "application/javascript")
	//字符过滤，防止XSS攻击
	callback := template.JSEscapeString(callbackFunc)
	_, err := ctx.Writer.Write([]byte(callback))
	if err != nil {
		return ctx
	}
	_, err = ctx.Writer.Write([]byte("("))
	if err != nil {
		return ctx
	}
	body, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}
	_, err = ctx.Writer.Write(body)
	if err != nil {
		return ctx
	}
	_, err = ctx.Writer.Write([]byte(")"))
	if err != nil {
		return ctx
	}
	return ctx
}

func (ctx *Context) IXml(obj interface{}) IResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.ISetHeader("Content-Type", "application/json")
	if _, err = ctx.Writer.Write(byt); err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	return ctx
}

func (ctx *Context) IHtml(file string, obj interface{}) IResponse {
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}
	if err = t.Execute(ctx.Writer, obj); err != nil {
		return ctx
	}
	ctx.ISetHeader("Content-Type", "application/html")
	return ctx
}

func (ctx *Context) IText(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values)
	ctx.ISetHeader("Content-Type", "application/text")
	_, _ = ctx.Writer.Write([]byte(out))
	return ctx
}

func (ctx *Context) IRedirect(path string) IResponse {
	http.Redirect(ctx.Writer, ctx.Request, path, http.StatusMovedPermanently)
	return ctx
}

func (ctx *Context) ISetHeader(key, val string) IResponse {
	ctx.Writer.Header().Add(key, val)
	return ctx
}

func (ctx *Context) ISetCookie(key, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
		SameSite: 1,
	})
	return ctx
}

func (ctx *Context) ISetStatus(code int) IResponse {
	ctx.Writer.WriteHeader(code)
	return ctx
}

func (ctx *Context) ISetOkStatus() IResponse {
	ctx.ISetStatus(http.StatusOK)
	return ctx
}
