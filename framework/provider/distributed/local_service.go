package distributed

import (
	"time"
	"webman/framework"
	"webman/framework/contract"
)

var _ contract.Distributed = (*LocalDistributedService)(nil)

type LocalDistributedService struct {
	container framework.Container
}

func NewLocalDistributedService(params ...interface{}) (interface{}, error) {
	return nil, nil
}

func (s *LocalDistributedService) Select(serviceName, appId string, holdTime time.Duration) (selectApp string, err error) {
	return "", nil
}
