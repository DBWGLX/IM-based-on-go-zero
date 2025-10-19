package logic

import (
	"context"

	"go-chat/user/internal/svc"
	"go-chat/user/user"
	"go-chat/user/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.Model.Insert(l.ctx, &model.User{
		Username: in.Username,
		Account:  in.Account,
		Password: in.Password,
	})

	return &user.RegisterResponse{}, nil
}
