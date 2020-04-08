package rpc

import (
	"context"
	msgpb "inn/pb/message"

	"inn/internal/message/service"
)

type messageRpc struct {
	ms *service.MessageService
}

func NewMessageRpc(ms *service.MessageService) *messageRpc {
	return &messageRpc{ms: ms}
}

func (mr *messageRpc) SendMsg(ctx context.Context, req *msgpb.SendMsgRequest, rsp *msgpb.MessageResponse) error {
	err := mr.ms.SendMessage(req.GetSenderUid(), req.GetRecipientUid(), req.GetContent(), int(req.GetMsgType()))
	return err
}

func (mr *messageRpc) QueryRecordFrom(ctx context.Context, req *msgpb.QueryRecordFromRequest, rsp *msgpb.MessagesResponse) error {
	messages, err := mr.ms.QueryRecordFrom(req.GetOwnerUid(), req.GetOtherUid(), req.GetFromMid(), req.GetCount())

	rsp.List = make([]*msgpb.MessageResponse, len(messages))

	for i, message := range messages {
		rsp.List[i] = &msgpb.MessageResponse{
			Mid:        message.Mid,
			Content:    message.Content,
			OwnerUid:   message.OwnerUid,
			Type:       int32(message.Type),
			OtherUid:   message.OtherUid,
			CreateTime: message.CreateTime,
			// OwnerUidAvatar: message.OwnerUidAvatar,
			// OtherUidAvatar: message.OtherUidAvatar,
			OwnerName: message.OwnerName,
			OtherName: message.OtherName,
		}
	}
	return err
}

func (mr *messageRpc) QueryNewerMsgFrom(ctx context.Context, req *msgpb.QueryNewerMsgFromRequest, rsp *msgpb.MessagesResponse) error {
	messages, err := mr.ms.QueryNewerMsgFrom(req.GetOwnerUid(), req.GetOtherUid(), req.GetFromMid())

	rsp.List = make([]*msgpb.MessageResponse, len(messages))

	for i, message := range messages {
		rsp.List[i] = &msgpb.MessageResponse{
			Mid:        message.Mid,
			Content:    message.Content,
			OwnerUid:   message.OwnerUid,
			Type:       int32(message.Type),
			OtherUid:   message.OtherUid,
			CreateTime: message.CreateTime,
			// OwnerUidAvatar: message.OwnerUidAvatar,
			// OtherUidAvatar: message.OtherUidAvatar,
			OwnerName: message.OwnerName,
			OtherName: message.OtherName,
		}
	}

	return err
}

func (mr *messageRpc) QueryContacts(ctx context.Context, req *msgpb.QueryContactsRequest, rsp *msgpb.MessageContactResponse) error {

	contactVO, _ := mr.ms.QueryContacts(req.GetOwnerUid())

	rsp.OwnerUid = contactVO.OwnerUid
	//rsp.OwnerAvatar  =  contactVO.OwnerAvatar
	rsp.OwnerName = contactVO.OwnerName
	rsp.TotalUnread = contactVO.TotalUnread
	rsp.ContactInfoList = make([]*msgpb.ContactInfo, len(contactVO.ContactInfos))
	for i, v := range contactVO.ContactInfos {
		rsp.ContactInfoList[i] = &msgpb.ContactInfo{
			OtherUid:    v.OtherUid,
			OtherName:   v.OtherName,
			OtherAvatar: v.OtherAvatar,
			Mid:         v.Mid,
			Type:        int32(v.Type),
			Content:     v.Content,
			ConvUnread:  v.ConvUnread,
			CreateTime:  v.CreateTime,
		}
	}

	return nil
}

func (mr *messageRpc) QueryTotalUnread(ctx context.Context, req *msgpb.QueryTotalUnreadRequest, rsp *msgpb.UnreadResponse) error {
	return nil
}
