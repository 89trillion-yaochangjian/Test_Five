package wsClient

import (
	"ChatService/internal/log"
	"ChatService/internal/model"
	"google.golang.org/protobuf/proto"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// 允许向对等方写入消息的时间。
	writeWait = 10 * time.Second

	// 允许从对等方读取下一个 pong 消息的时间。
	pongWait = 60 * time.Second

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

//读取

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Info.Println(err)
			}
			break
		}
		msg := &model.ChatRequest{}
		proto.Unmarshal(message, msg)
		if msg.Type == model.ExitType {
			err := ExitType(UserList, msg, c)
			if err != nil {
				break
			}
		} else if msg.Type == model.TalkType {
			log.Info.Print(model.TalkLog, msg.Content)
			err := TalkType(UserList, msg, c)
			if err != nil {
				break
			}
		} else if msg.Type == model.UserListType {
			//读取用户列表
			err := UserListType(msg, c)
			if err != nil {
				break
			}
		} else {
			newMsg, _ := proto.Marshal(msg)
			c.Send <- newMsg
		}
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

//Talk 类型消息处理

func TalkType(UserList map[string]string, msg *model.ChatRequest, c *Client) error {
	UserList[msg.UserName] = msg.UserName
	msg.UserList = UserList
	newMsg, err := proto.Marshal(msg)
	if err != nil {
		log.Error.Println(err)
		return err
	}
	c.Hub.Broadcast <- newMsg
	return err
}

//Exit 类型消息处理

func ExitType(UserList map[string]string, msg *model.ChatRequest, c *Client) error {
	delete(UserList, msg.UserName)
	msg.UserList = UserList
	newMsg, err := proto.Marshal(msg)
	if err != nil {
		log.Error.Println(err)
		return err
	}
	c.Hub.Broadcast <- newMsg
	c.Hub.Unregister <- c
	return err
}

//UserList 类型消息处理

func UserListType(msg *model.ChatRequest, c *Client) error {
	var userList string
	if UserList != nil {
		for _, value := range UserList {
			userList += value + ","
		}
	}
	msg.Content = userList
	newMsg, err := proto.Marshal(msg)
	if err != nil {
		log.Error.Println(err)
		return err
	}
	c.Hub.Broadcast <- newMsg
	return err
}
