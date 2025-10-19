package logic

import (
	"context"

	"go-chat/im/im"
	"go-chat/im/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryLogic {
	return &GetHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetHistoryLogic) GetHistory(in *im.GetHistoryRequest) (*im.GetHistoryResponse, error) {
	// todo: add your logic here and delete this line

	return &im.GetHistoryResponse{}, nil
}
