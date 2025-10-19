// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package websocketHandle

import (
	"net/http"

	"go-chat/api/internal/logic/websocketLogic"
	"go-chat/api/internal/svc"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		//允许跨域请求
		return true
	},
}

// WebSocket 连接
func WsChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//1. 从 JWT token 中获取 userId
		//【假设 userId 已经通过中间件放入 r.Context()
		userId, ok := r.Context().Value("userId").(string) //【类型断言 【panic
		if userId == "" || !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		//2.升级ws
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logx.Error("upgrade error:", err)
			return
		}

		//3.注册
		client := &svc.Client{
			Conn:   conn,
			UserId: userId,
		}

		//4.启动
		//接收用户消息
		go websocketLogic.ReadPump(client, svcCtx.IMRpc)
		//转发用户消息
		//go writePump()

	}
}
