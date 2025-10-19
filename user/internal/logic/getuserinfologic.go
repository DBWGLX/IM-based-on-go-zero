package logic

import (
	"context"

	"go-chat/user/internal/svc"
	"go-chat/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	ret := user.UserInfoResponse{}

	res, err := l.svcCtx.Model.FindOneByAccount(l.ctx, in.Account)
	if err != nil {
		return nil, err
	}
	ret.Account = res.Account
	ret.Username = res.Username

	//getAvatar

	return &ret, nil
}
