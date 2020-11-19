package svc

import (
	"github.com/lqlspace/microservices/src/basic/cmd/api/config"
)

type (
	ServiceContext struct {
	}
)

func NewServiceContext(c *config.Config)  *ServiceContext {
	return &ServiceContext{
	}
}
