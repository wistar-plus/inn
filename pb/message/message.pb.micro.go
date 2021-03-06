// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Message service

type MessageService interface {
	SendMsg(ctx context.Context, in *SendMsgRequest, opts ...client.CallOption) (*MessageResponse, error)
	QueryRecordFrom(ctx context.Context, in *QueryRecordFromRequest, opts ...client.CallOption) (*MessagesResponse, error)
	QueryNewerMsgFrom(ctx context.Context, in *QueryNewerMsgFromRequest, opts ...client.CallOption) (*MessagesResponse, error)
	QueryContacts(ctx context.Context, in *QueryContactsRequest, opts ...client.CallOption) (*MessageContactResponse, error)
	QueryTotalUnread(ctx context.Context, in *QueryTotalUnreadRequest, opts ...client.CallOption) (*UnreadResponse, error)
}

type messageService struct {
	c    client.Client
	name string
}

func NewMessageService(name string, c client.Client) MessageService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "message"
	}
	return &messageService{
		c:    c,
		name: name,
	}
}

func (c *messageService) SendMsg(ctx context.Context, in *SendMsgRequest, opts ...client.CallOption) (*MessageResponse, error) {
	req := c.c.NewRequest(c.name, "Message.SendMsg", in)
	out := new(MessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) QueryRecordFrom(ctx context.Context, in *QueryRecordFromRequest, opts ...client.CallOption) (*MessagesResponse, error) {
	req := c.c.NewRequest(c.name, "Message.QueryRecordFrom", in)
	out := new(MessagesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) QueryNewerMsgFrom(ctx context.Context, in *QueryNewerMsgFromRequest, opts ...client.CallOption) (*MessagesResponse, error) {
	req := c.c.NewRequest(c.name, "Message.QueryNewerMsgFrom", in)
	out := new(MessagesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) QueryContacts(ctx context.Context, in *QueryContactsRequest, opts ...client.CallOption) (*MessageContactResponse, error) {
	req := c.c.NewRequest(c.name, "Message.QueryContacts", in)
	out := new(MessageContactResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) QueryTotalUnread(ctx context.Context, in *QueryTotalUnreadRequest, opts ...client.CallOption) (*UnreadResponse, error) {
	req := c.c.NewRequest(c.name, "Message.QueryTotalUnread", in)
	out := new(UnreadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Message service

type MessageHandler interface {
	SendMsg(context.Context, *SendMsgRequest, *MessageResponse) error
	QueryRecordFrom(context.Context, *QueryRecordFromRequest, *MessagesResponse) error
	QueryNewerMsgFrom(context.Context, *QueryNewerMsgFromRequest, *MessagesResponse) error
	QueryContacts(context.Context, *QueryContactsRequest, *MessageContactResponse) error
	QueryTotalUnread(context.Context, *QueryTotalUnreadRequest, *UnreadResponse) error
}

func RegisterMessageHandler(s server.Server, hdlr MessageHandler, opts ...server.HandlerOption) error {
	type message interface {
		SendMsg(ctx context.Context, in *SendMsgRequest, out *MessageResponse) error
		QueryRecordFrom(ctx context.Context, in *QueryRecordFromRequest, out *MessagesResponse) error
		QueryNewerMsgFrom(ctx context.Context, in *QueryNewerMsgFromRequest, out *MessagesResponse) error
		QueryContacts(ctx context.Context, in *QueryContactsRequest, out *MessageContactResponse) error
		QueryTotalUnread(ctx context.Context, in *QueryTotalUnreadRequest, out *UnreadResponse) error
	}
	type Message struct {
		message
	}
	h := &messageHandler{hdlr}
	return s.Handle(s.NewHandler(&Message{h}, opts...))
}

type messageHandler struct {
	MessageHandler
}

func (h *messageHandler) SendMsg(ctx context.Context, in *SendMsgRequest, out *MessageResponse) error {
	return h.MessageHandler.SendMsg(ctx, in, out)
}

func (h *messageHandler) QueryRecordFrom(ctx context.Context, in *QueryRecordFromRequest, out *MessagesResponse) error {
	return h.MessageHandler.QueryRecordFrom(ctx, in, out)
}

func (h *messageHandler) QueryNewerMsgFrom(ctx context.Context, in *QueryNewerMsgFromRequest, out *MessagesResponse) error {
	return h.MessageHandler.QueryNewerMsgFrom(ctx, in, out)
}

func (h *messageHandler) QueryContacts(ctx context.Context, in *QueryContactsRequest, out *MessageContactResponse) error {
	return h.MessageHandler.QueryContacts(ctx, in, out)
}

func (h *messageHandler) QueryTotalUnread(ctx context.Context, in *QueryTotalUnreadRequest, out *UnreadResponse) error {
	return h.MessageHandler.QueryTotalUnread(ctx, in, out)
}
