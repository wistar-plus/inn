package controller

import (
	"inn/internal/user/model"
	msgpb "inn/pb/message"
	userpb "inn/pb/user"
	"inn/pkg/e"
	"inn/pkg/gintool"
	"log"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSrv    userpb.UserService
	messageSrv msgpb.MessageService
}

func NewUserController(userSrv userpb.UserService, messageSrv msgpb.MessageService) *UserController {
	return &UserController{userSrv: userSrv, messageSrv: messageSrv}
}

func (uc *UserController) Register(c *gin.Context) {
	var param *model.UserRegisterParam

	err := c.ShouldBindJSON(&param)
	if err != nil {
		gintool.ResError(c, e.INVALID_PARAMS, err)
		return
	}

	req := &userpb.RegisterRequest{
		Email:    param.Email,
		Pwd:      param.Pwd,
		NickName: param.NickName,
	}

	ctx, ok := gintool.ContextWithSpan(c)
	if !ok {
		log.Println("get context err")
	}
	_, err = uc.userSrv.Register(ctx, req)
	if err != nil {
		gintool.ResError(c, e.Fail, err)
		return
	}

	gintool.ResSuccess(c, nil)
}
func (uc *UserController) Login(c *gin.Context) {
	var param *model.UserLoginParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		gintool.ResError(c, e.INVALID_PARAMS, err)
		return
	}

	req := &userpb.LoginRequest{
		Email: param.Email,
		Pwd:   param.Pwd,
	}

	ctx, ok := gintool.ContextWithSpan(c)
	if !ok {
		log.Println("get context err")
	}
	user, err := uc.userSrv.Login(ctx, req)
	if err != nil {
		gintool.ResError(c, e.Fail, err)
		return
	}

	rsp, err := uc.userSrv.GetAllUsersExcept(ctx, &userpb.UserIdRequest{Uid: user.GetUid()})
	if err != nil {
		gintool.ResError(c, e.Fail, err)
	}

	contacts, err := uc.messageSrv.QueryContacts(ctx, &msgpb.QueryContactsRequest{OwnerUid: user.GetUid()})
	if err != nil {
		gintool.ResError(c, e.Fail, err)
	}

	gintool.ResSuccess(c, map[string]interface{}{
		"loginUser":   user,
		"userList":    rsp.Users,
		"contactList": contacts,
	})
}
