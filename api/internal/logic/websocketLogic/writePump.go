package websocketLogic

import (
	"go-chat/api/internal/svc"
	"go-chat/im/imclient"
)

// 写入口
func WritePump(client *svc.Client, imRpc imclient.IM) {
	defer client.Conn.Close()
}
