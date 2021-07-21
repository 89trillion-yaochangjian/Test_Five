package ctrl

import (
	"ChatService/internal/config"
	"ChatService/internal/model"
	"google.golang.org/protobuf/proto"
)

//Talk 类型消息处理

func TalkType(UserList map[string]string, msg *model.ChatRequest) ([]byte, error) {
	UserList[msg.UserName] = msg.UserName
	msg.UserList = UserList
	newMsg, err := proto.Marshal(msg)
	if err != nil {
		config.Error.Println(err)
		return newMsg, err
	}
	return newMsg, err
}

//Exit 类型消息处理

func ExitType(UserList map[string]string, msg *model.ChatRequest) ([]byte, error) {
	delete(UserList, msg.UserName)
	msg.UserList = UserList
	msg.Content = "exit"
	newMsg, err := proto.Marshal(msg)
	if err != nil {
		config.Error.Println(err)
		return newMsg, err
	}
	return newMsg, err
}

//UserList 类型消息处理

func UserListType(msg *model.ChatRequest, UserList map[string]string) ([]byte, error) {
	var userList string
	if UserList != nil {
		for _, value := range UserList {
			userList += value + ","
		}
	}
	msg.Content = userList
	newMsg, err := proto.Marshal(msg)
	if err != nil {
		config.Error.Println(err)
		return newMsg, err
	}
	return newMsg, err
}
