package wsClient

//抽象连接器
type Hub struct {
	// 注册客户
	Clients map[*Client]bool
	// 来自客户端的发送消息
	Broadcast chan []byte
	// 注册来自客户端的请求
	Register chan *Client
	//取消注册来自客户端的请求。
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			//注册标志
			h.Clients[client] = true
		case client := <-h.Unregister:
			//注销
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			//处理数据
			//client具体每一个链接
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

