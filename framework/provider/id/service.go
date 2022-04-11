package id

import "github.com/rs/xid"

type IDService struct {
}

func NewIDService() *IDService {
	return &IDService{}
}

func NewID() string {
	return xid.New().String()
}
