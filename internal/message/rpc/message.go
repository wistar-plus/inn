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
	rsp.Mid = 1
	rsp.Content = "hello"
	return err
}

func (mr *messageRpc) QueryConversationMsg(ctx context.Context, req *msgpb.QueryConversationMsgRequest, rsp *msgpb.MessageResponse) error {
	return nil
}

func (mr *messageRpc) QueryNewerMsgFrom(ctx context.Context, req *msgpb.QueryNewerMsgFromRequest, rsp *msgpb.MessageResponse) error {
	return nil
}

func (mr *messageRpc) QueryContacts(ctx context.Context, req *msgpb.QueryContactsRequest, rsp *msgpb.MessageContactResponse) error {
	return nil
}

func (mr *messageRpc) QueryTotalUnread(ctx context.Context, req *msgpb.QueryTotalUnreadRequest, rsp *msgpb.UnreadResponse) error {
	return nil
}
