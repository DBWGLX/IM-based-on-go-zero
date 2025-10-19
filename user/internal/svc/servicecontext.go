package svc

import (
	"go-chat/user/internal/config"
	"go-chat/user/user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	Model model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Model: model.NewUserModel(sqlx.NewMysql(c.DataSource)),
	}
}
