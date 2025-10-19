package websocketLogic

import (
	"encoding/json"
	"go-chat/api/internal/svc"
	"go-chat/api/internal/types"
	"go-chat/im/imclient"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

// 读入口
func ReadPump(client *svc.Client, imRpc imclient.IM) {
	defer func() {
		svc.GlobalHub.DoUnregister(client)
		client.Conn.Close()
	}()

	for {
		// 1. 读取消息
		_, msgBytes, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logx.Error("read error:", err)
			}
			break
		}

		// 2.解析
		var msg types.Message
		if err := json.Unmarshal(msgBytes, &msg); err != nil {
			logx.Error("invalid message format:", err)
			continue
		}

		// 3.存储

		// 4.在线发送

		//
	}
}
