package kernel

import (
	"net/http"
	"webman/framework/gin"
)

type WmKernelService struct {
	engine *gin.Engine
}

func NewWmKernelService(params ...interface{}) (interface{}, error) {
	engine := params[0].(*gin.Engine)
	return &WmKernelService{engine: engine}, nil
}

func (s *WmKernelService) HttpEngine() http.Handler {
	return s.engine
}
