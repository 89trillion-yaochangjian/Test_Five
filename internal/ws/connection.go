package ws

import (
	"ChatService/internal/config"
	"ChatService/internal/ctrl"
	"ChatService/internal/model"
	"ChatService/internal/router"
	"google.golang.org/protobuf/proto"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// 允许向对等方写入消息的时间。
	writeWait = 10 * time.Second

	// 允许从对等方读取下一个 pong 消息的时间。
	pongWait = 10 * time.Second

	// 在此期间向对等方发送 ping。 必须小于 pongWait。
	pingPeriod = (pongWait * 9) / 10

	// 对等方允许的最大消息大小。
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Client struct {
	Hub *Hub
	// websocket 连接器
	Conn *websocket.Conn
	//消息的缓冲通道。
	Send        chan []byte
	ChatRequest model.ChatRequest
}

var UserList = make(map[string]string)
var msg = &model.ChatRequest{}

//读取

func (c *Client) ReadPump() {
	defer func() {
		newMsg, _ := ctrl.ExitType(UserList, msg)
		c.Hub.Broadcast <- newMsg
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				config.Info.Println(err)
			}
			break
		}
		newMsg := router.Type(UserList, message)
		c.Hub.Broadcast <- newMsg
		proto.Unmarshal(message, msg)
	}
}

// 执行写入

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		//从管道中读取数据
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
