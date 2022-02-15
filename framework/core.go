package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	//router map[string]map[string]ControllerHandler
	router map[string]*Tree
	//存放中间件
	middlewares []ControllerHandler
}

func NewCore() *Core {
	////定义二级map
	//getRouter := map[string]ControllerHandler{}
	//postRouter := map[string]ControllerHandler{}
	//putRouter := map[string]ControllerHandler{}
	//deleteRouter := map[string]ControllerHandler{}
	////定义一级map
	//router := map[string]map[string]ControllerHandler{}
	//router["GET"] = getRouter
	//router["POST"] = postRouter
	//router["PUT"] = putRouter
	//router["DELETE"] = deleteRouter
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	return &Core{router: router}
}

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	//c.router["GET"][upperUrl] = handler
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	//c.router["POST"][upperUrl] = handler
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	//c.router["Put"][upperUrl] = handler
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	//c.router["Delete"][upperUrl] = handler
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

func (c *Core) FindRouterByRequest(r *http.Request) []ControllerHandler {
	uri := r.URL.Path
	method := r.Method

	upperUri := strings.ToUpper(uri)
	upperMethod := strings.ToUpper(method)

	if methodHandlers, ok := c.router[upperMethod]; ok {
		//if handler, ok := methodHandlers[upperUri]; ok {
		//	return handler
		//}
		return methodHandlers.FindHandler(upperUri)
	}
	return nil
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(r, w)

	handlers := c.FindRouterByRequest(r)
	if handlers == nil {
		ctx.Json(404, "not found")
		return
	}
	ctx.SetHandlers(handlers)

	if err := ctx.Next(); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}
