// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"go-chat/api/internal/config"
	"go-chat/im/imclient"
	"go-chat/user/userclient"
)

type ServiceContext struct {
	Config config.Config

	IMRpc   imclient.IM
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
