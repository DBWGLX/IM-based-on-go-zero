// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package websocket

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-chat/api/internal/logic/websocket"
	"go-chat/api/internal/svc"
	"go-chat/api/internal/types"
)

// WebSocket 连接
func WsChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := websocket.NewWsChatLogic(r.Context(), svcCtx)
		err := l.WsChat(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
