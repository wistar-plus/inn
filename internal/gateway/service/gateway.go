package service

import (
	"context"
	"fmt"
	"inn/internal/gateway/model"
	"inn/internal/gateway/repository"
	msgpb "inn/pb/message"

	"github.com/spf13/viper"
)

type GateWayService struct {
	userConnRepo  repository.IUserConnRepository
	userTopicRepo repository.IUserTopicRepository
	messageSrvice msgpb.MessageService
}

func NewGateWayService(userConnRepo repository.IUserConnRepository, userTopicRepo repository.IUserTopicRepository, messageSrvice msgpb.MessageService) *GateWayService {
	return &GateWayService{userConnRepo: userConnRepo, userTopicRepo: userTopicRepo, messageSrvice: messageSrvice}
}

func (gs *GateWayService) StoreConn(uid uint64, conn *model.Conn) {
	gs.userConnRepo.Add(uid, conn)
	gs.userTopicRepo.Add(uid, viper.GetString("topic"))
	// if c,ok := gs.userConnRepo.Get(uid) {
	// 	fmt.Println(c.Uid)
	// }
}

func (gs *GateWayService) SendMessage(ctx context.Context, senderUid, recipientUid uint64, content string, msgType int32) {
	msg := &msgpb.SendMsgRequest{
		SenderUid:    senderUid,
		RecipientUid: recipientUid,
		Content:      content,
		MsgType:      msgType,
	}
	fmt.Println("发送消息给 msgSrv: ", msg)

	_, err := gs.messageSrvice.SendMsg(ctx, msg)
	if err != nil {
		fmt.Println(err)
	}
}

func (gs *GateWayService) GetRecordFromMid(ctx context.Context, ownerUid, otherUid, mid uint64, count int64) *msgpb.MessagesResponse {
	req := &msgpb.QueryRecordFromRequest{
		OwnerUid: ownerUid,
		OtherUid: otherUid,
		FromMid:  mid,
		Count:    count,
	}

	rsp, err := gs.messageSrvice.QueryRecordFrom(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	return rsp
}
