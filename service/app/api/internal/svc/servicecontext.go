package svc

import "tinkgo/service/app/api/internal/config"

type ServiceContext struct {
	Config *config.Config
}

func NewServiceContext(config *config.Config) (srvCtx *ServiceContext, err error) {
	srvCtx = &ServiceContext{
		Config: config,
	}
	return
}
