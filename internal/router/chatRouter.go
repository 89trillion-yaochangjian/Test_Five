package router

import (
	"ChatService/internal/config"
	"ChatService/internal/ctrl"
	"ChatService/internal/model"
	"google.golang.org/protobuf/proto"
)

func Type(UserList map[string]string, message []byte) []byte {
	msg := &model.ChatRequest{}
	proto.Unmarshal(message, msg)
	if msg.Type == model.ExitType {
		newMsg, err := ctrl.ExitType(UserList, msg)
		if err != nil {
			config.Error.Println(err)
			return newMsg
		}
		return newMsg
	} else if msg.Type == model.UserListType {
		//读取用户列表
		newMsg, err := ctrl.UserListType(msg, UserList)
		if err != nil {
			config.Error.Println(err)
			return newMsg
		}
		return newMsg
	} else {
		config.Info.Print(model.TalkLog, msg.Content)
		newMsg, err := ctrl.TalkType(UserList, msg)
		if err != nil {
			config.Error.Println(err)
			return newMsg
		}
		return newMsg
	}
}
