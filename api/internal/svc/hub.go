package svc

import (
	"sync"

	"github.com/gorilla/websocket"
)

// 一个ws连接
type Client struct {
	Conn   *websocket.Conn
	UserId string
}

// 管理连接
type Hub struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	lock       sync.RWMutex
}

var GlobalHub *Hub

func InitHub() {
	GlobalHub = &Hub{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go GlobalHub.run()
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.DoRegister(client)
		case client := <-h.unregister:
			h.DoUnregister(client)
		}
	}
}

func (h *Hub) DoRegister(client *Client) {
	h.lock.Lock()
	defer h.lock.Unlock()

	// 注册，如果已有连接，先关闭旧的
	if oldClient, ok := h.clients[client.UserId]; ok {
		oldClient.Conn.Close()
	}
	h.clients[client.UserId] = client
}

func (h *Hub) DoUnregister(client *Client) {
	h.lock.Lock()
	defer h.lock.Unlock()

	if _, ok := h.clients[client.UserId]; ok {
		delete(h.clients, client.UserId)
		//close(client.send) // 假设 Client 有 send channel
	}
}

// 获取连接
func (h *Hub) GetClient(userId string) *Client {
	h.lock.RLock()
	defer h.lock.RUnlock()

	if client, ok := h.clients[userId]; ok {
		return client
	}
	return nil
}

//读

//写
