// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package websocketLogic

import (
	"context"

	"go-chat/api/internal/svc"
	"go-chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WsChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// WebSocket 连接
func NewWsChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsChatLogic {
	return &WsChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsChatLogic) WsChat(req *types.WsReq) error {
	// todo: add your logic here and delete this line

	return nil
}
