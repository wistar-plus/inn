package service

import (
	"encoding/json"
	"fmt"
	"inn/internal/message/model"
	"inn/internal/message/repository"

	"github.com/micro/go-micro/broker"
)

type MessageService struct {
	//repo repository.IMessageRepository
	pubsub        broker.Broker
	userTopicRepo repository.IUserTopicRepository
}

func NewMessageService(pubsub broker.Broker, userTopicRepo repository.IUserTopicRepository) *MessageService {
	return &MessageService{
		//repo: repo,
		pubsub:        pubsub,
		userTopicRepo: userTopicRepo,
	}
}

func (ms *MessageService) SendMessage(senderUid, recipientUid uint64, content string, msgType int) error {
	/**存内容*/
	messageContent := new(model.MessageContent)
	messageContent.SenderId = senderUid
	messageContent.RecipientId = recipientUid
	messageContent.Content = content
	messageContent.MsgType = msgType
	// messageContent = contentRepository.Insert(messageContent);
	// mid := messageContent.Mid

	/**存发件人的发件箱*/

	/**存收件人的收件箱*/

	/**更新发件人的最近联系人 */

	/**更新收件人的最近联系人 */

	/**更未读更新 */

	/** 待推送消息发布到mq */

	topic := ms.userTopicRepo.Get(messageContent.RecipientId)

	fmt.Printf("topic: %s, msg: %v\n", topic, messageContent)

	body, err := json.Marshal(messageContent)
	if err != nil {
		return err
	}

	msg := &broker.Message{
		Header: map[string]string{
			"id": "",
		},
		Body: body,
	}
	if err := ms.pubsub.Publish(topic, msg); err != nil {
		fmt.Errorf("publish message failed: %v\n", err)
		return err
	}
	return nil
	//return messageVO,nil
}

func (ms *MessageService) QueryConversationMsg() error {
	return nil
}

func (ms *MessageService) QueryNewerMsgFrom() error {
	return nil
}

func (m *MessageService) QueryContacts() error {
	return nil
}

func (m *MessageService) QueryTotalUnread() error {
	return nil
}
