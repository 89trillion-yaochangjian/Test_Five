package service

import (
	"ChatService/internal/config"
	"ChatService/internal/model"
	"ChatService/internal/ws"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// serveWs 处理来自对等方的 websocket 请求。

func ServeWs(hub *ws.Hub, w http.ResponseWriter, r *http.Request) {

	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		config.Error.Println(err)
		return
	}
	username := r.Header.Get("username")
	client := &ws.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256), ChatRequest: model.ChatRequest{}}
	hub.Register <- client
	client.ChatRequest.UserName = username
	// 允许通过完成所有工作来收集调用者引用的内存
	// 新的协程。
	go client.WritePump()
	go client.ReadPump()
}
