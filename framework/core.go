package framework

import (
	"log"
	"net/http"
)

type Core struct {
	routers map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{routers: map[string]ControllerHandler{}}
}

func (c *Core) Set(url string, handler ControllerHandler) {
	c.routers[url] = handler
}

func (c *Core) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("core.ServeHTTP")
	ctx := NewContext(r, w)

	router := c.routers["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")

	if err := router(ctx); err != nil {
		return
	}
}
