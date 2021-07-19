package service

import (
	"ChatService/internal/log"
	"ChatService/internal/wsClient"
	"github.com/gorilla/websocket"
	"net/http"
)
var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
// serveWs 处理来自对等方的 websocket 请求。

func ServeWs(hub *wsClient.Hub, w http.ResponseWriter, r *http.Request) {

	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error.Println(err)
		return
	}
	client := &wsClient.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	hub.Register <- client

	// 允许通过完成所有工作来收集调用者引用的内存
	// 新的协程。
	go client.WritePump()
	go client.ReadPump()
}

//func ServeHome(w http.ResponseWriter, r *http.Request) {
//	log.Println(r.URL)
//	if r.URL.Path != "/" {
//		http.Error(w, "Not found", http.StatusNotFound)
//		return
//	}
//	if r.Method != "GET" {
//		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//		return
//	}
//	http.ServeFile(w, r, "home.html")
//}