package service

import (
	"context"
	"encoding/json"
	"fmt"
	"inn/internal/message/model"
	"inn/internal/message/repository"
	"inn/pkg/e"

	userpb "inn/pb/user"

	"github.com/micro/go-micro/broker"
)

type MessageService struct {
	pubsub        broker.Broker
	userTopicRepo repository.IUserTopicRepository
	contentRepo   repository.IContentRepository
	relationRepo  repository.IRelationRepository
	contactRepo   repository.IContactRepository
	unreadRepo    repository.IUnreadRepository
	userService   userpb.UserService
}

func NewMessageService(pubsub broker.Broker, userTopicRepo repository.IUserTopicRepository,
	relationRepo repository.IRelationRepository, contentRepo repository.IContentRepository,
	contactRepo repository.IContactRepository, unreadRepo repository.IUnreadRepository, userService userpb.UserService) *MessageService {
	return &MessageService{
		pubsub:        pubsub,
		userTopicRepo: userTopicRepo,
		contentRepo:   contentRepo,
		relationRepo:  relationRepo,
		contactRepo:   contactRepo,
		unreadRepo:    unreadRepo,
		userService:   userService,
	}
}

func (ms *MessageService) SendMessage(senderUid, recipientUid uint64, content string, msgType int) error {
	/**存内容*/
	messageContent := new(model.MessageContent)
	messageContent.SenderId = senderUid
	messageContent.RecipientId = recipientUid
	messageContent.Content = content
	messageContent.MsgType = msgType
	mid, err := ms.contentRepo.Insert(messageContent)
	if err != nil {
		fmt.Println(err)
		return e.ERROR_DBERROR
	}
	//mid := messageContent.Mid

	/**存发件人的发件箱*/
	messageRelationSender := new(model.MessageRelation)
	messageRelationSender.Mid = mid
	messageRelationSender.OwnerUid = senderUid
	messageRelationSender.OtherUid = recipientUid
	messageRelationSender.Type = 0
	ms.relationRepo.Insert(messageRelationSender)

	/**存收件人的收件箱*/
	messageRelationRecipient := new(model.MessageRelation)
	messageRelationRecipient.Mid = mid
	messageRelationRecipient.OwnerUid = recipientUid
	messageRelationRecipient.OtherUid = senderUid
	messageRelationRecipient.Type = 1
	ms.relationRepo.Insert(messageRelationRecipient)

	/**更新发件人的最近联系人 */
	messageContactSender := ms.contactRepo.FindOne(senderUid, recipientUid)
	if messageContactSender != nil {
		messageContactSender.Mid = mid
	} else {
		messageContactSender = new(model.MessageContact)
		messageContactSender.OwnerUid = senderUid
		messageContactSender.OtherUid = recipientUid
		messageContactSender.Mid = mid
		messageContactSender.Type = 0
	}
	ms.contactRepo.Save(messageContactSender)

	/**更新收件人的最近联系人 */
	messageContactRecipient := ms.contactRepo.FindOne(recipientUid, senderUid)
	if messageContactRecipient != nil {
		messageContactRecipient.Mid = mid
	} else {
		messageContactRecipient = new(model.MessageContact)
		messageContactRecipient.OwnerUid = recipientUid
		messageContactRecipient.OtherUid = senderUid
		messageContactRecipient.Mid = mid
		messageContactRecipient.Type = 0
	}
	ms.contactRepo.Save(messageContactRecipient)

	/**更未读更新 */

	ms.unreadRepo.IncrementTotalUnreadBy(recipientUid, 1)
	ms.unreadRepo.IncrementUnreadBy(recipientUid, senderUid, 1)

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
}

func (ms *MessageService) QueryRecordFrom(ownerUid, otherUid, mid uint64, count int64) ([]*model.MessageVO, error) {
	relationList, err := ms.relationRepo.FindAllByOwnerUidAndOtherUidAndMidIsLessThanOrderByMidDescLimit(ownerUid, otherUid, mid, count)
	if err != nil {
		return nil, err
	}
	return ms.composeMessage(relationList, ownerUid, otherUid)
}

func (ms *MessageService) QueryNewerMsgFrom(ownerUid, otherUid, mid uint64) ([]*model.MessageVO, error) {
	relationList, err := ms.relationRepo.FindAllByOwnerUidAndOtherUidAndMidIsGreaterThanOrderByMidAsc(ownerUid, otherUid, mid)
	if err != nil {
		return nil, err
	}
	return ms.composeMessage(relationList, ownerUid, otherUid)
}

func (ms *MessageService) composeMessage(relationList []*model.MessageRelation, ownerUid, otherUid uint64) ([]*model.MessageVO, error) {
	/** 先拼接消息索引和内容 */
	msgList := make([]*model.MessageVO, len(relationList))
	self, _ := ms.userService.GetUserById(context.TODO(), &userpb.UserIdRequest{Uid: ownerUid})
	other, _ := ms.userService.GetUserById(context.TODO(), &userpb.UserIdRequest{Uid: otherUid})
	for i, relation := range relationList {
		mid := relation.Mid
		content, _ := ms.contentRepo.FindById(mid)
		if content != nil {
			msgList[i] = &model.MessageVO{
				Mid:        mid,
				Content:    content.Content,
				OwnerUid:   relation.OwnerUid,
				Type:       relation.Type,
				OtherUid:   relation.OtherUid,
				CreateTime: relation.CreatedAt.Unix(),
				// OwnerUidAvatar : ,
				// OtherUidAvatar : ,
				OwnerName: self.NickName,
				OtherName: other.NickName,
			}
		}
	}

	/** 再变更未读 */
	convUnread, _ := ms.unreadRepo.GetUnread(ownerUid, otherUid)
	ms.unreadRepo.DelUnread(ownerUid, otherUid)
	afterCleanUnread, _ := ms.unreadRepo.IncrementTotalUnreadBy(ownerUid, -convUnread)
	if afterCleanUnread <= 0 {
		ms.unreadRepo.DelTotalUnread(ownerUid)
	}

	return msgList, nil
}

func (ms *MessageService) QueryContacts(ownerUid uint64) (*model.MessageContactVO, error) {
	contacts, _ := ms.contactRepo.FindMessageContactsByOwnerUidOrderByMidDesc(ownerUid)
	if contacts != nil {
		user, _ := ms.userService.GetUserById(context.TODO(), &userpb.UserIdRequest{Uid: ownerUid})
		totalUnread, _ := ms.unreadRepo.GetTotalUnread(user.GetUid())
		contactVO := &model.MessageContactVO{
			OwnerUid:     user.GetUid(),
			OwnerName:    user.GetNickName(),
			TotalUnread:  totalUnread,
			ContactInfos: make([]*model.ContactInfo, len(contacts)),
		}
		for i, contact := range contacts {
			mid := contact.Mid
			content, _ := ms.contentRepo.FindById(mid)
			otherUser, _ := ms.userService.GetUserById(context.TODO(), &userpb.UserIdRequest{Uid: contact.OtherUid})
			if contact != nil {
				convUnread, _ := ms.unreadRepo.GetUnread(user.GetUid(), otherUser.GetUid())
				contactVO.ContactInfos[i] = &model.ContactInfo{
					OtherUid:  otherUser.GetUid(),
					OtherName: otherUser.GetNickName(),
					//OtherAvatar string
					Mid:        mid,
					Type:       contact.Type,
					Content:    content.Content,
					ConvUnread: convUnread,
					CreateTime: contact.CreatedAt.Unix(),
				}
			}
		}
		return contactVO, nil
	}

	return nil, nil
}

func (ms *MessageService) QueryTotalUnread(uid uint64) (int64, error) {
	return ms.unreadRepo.GetTotalUnread(uid)
}
