package subscriber

import (
	"encoding/json"
	"errors"
	"fmt"
	"inn/internal/gateway/repository"
	"inn/internal/message/model"

	"github.com/micro/go-micro/broker"
)

type MessageSubscriber struct {
	userConnRepo repository.IUserConnRepository
}

func NewMessageSubscriber(userConnRepo repository.IUserConnRepository) *MessageSubscriber {
	return &MessageSubscriber{userConnRepo: userConnRepo}
}

func (msub *MessageSubscriber) Handler(pub broker.Event) error {
	var msgContent *model.MessageContent
	if err := json.Unmarshal(pub.Message().Body, &msgContent); err != nil {
		fmt.Errorf("process message failed: %v\n", err)
		return err
	}

	fmt.Println("mq推送消息：", msgContent)

	conn, ok := msub.userConnRepo.Get(msgContent.RecipientId)
	if !ok {
		return errors.New("conn no exist")
	}
	rsp := map[string]interface{}{
		"type": 3,
		"data": msgContent,
	}

	conn.Wsconn.WriteJSON(rsp)

	return nil
}
