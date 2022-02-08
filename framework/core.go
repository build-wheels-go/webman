package framework

import (
	"log"
	"net/http"
	"strings"
)

type Core struct {
	//router map[string]map[string]ControllerHandler
	router map[string]*Tree
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

func (c *Core) Get(url string, handler ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	//c.router["GET"][upperUrl] = handler
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handler ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	//c.router["POST"][upperUrl] = handler
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handler ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	//c.router["Put"][upperUrl] = handler
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	//upperUrl := strings.ToUpper(url)
	//c.router["Delete"][upperUrl] = handler
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

func (c *Core) FindRouterByRequest(r *http.Request) ControllerHandler {
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

	router := c.FindRouterByRequest(r)
	if router == nil {
		ctx.Json(404, "not found")
		return
	}

	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}
