package logic

import (
	"context"

	"go-chat/im/im"
	"go-chat/im/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStoreMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreMessageLogic {
	return &StoreMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StoreMessageLogic) StoreMessage(in *im.StoreMessageRequest) (*im.StoreMessageResponse, error) {
	// todo: add your logic here and delete this line

	return &im.StoreMessageResponse{}, nil
}
