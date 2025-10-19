// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"go-chat/api/internal/config"
	"go-chat/im/imclient"
	"go-chat/im/imservice"
	"go-chat/user/userclient"
	"go-chat/user/userservice"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	IMRpc   imclient.IM
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	//非must
	imClient, err := zrpc.NewClient(c.IMRpc)
	if err != nil {
		logx.Errorf("user rpc unavailable: %v", err)
	}

	return &ServiceContext{
		Config: c,

		IMRpc:   imservice.NewIMService(imClient),
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
	}
}
