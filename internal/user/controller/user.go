package controller

import (
	"inn/internal/user/model"
	"inn/internal/user/service"
	"inn/pkg/e"
	"inn/pkg/gintool"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	us *service.UserService
}

func NewUserController(us *service.UserService) *UserController {
	return &UserController{us: us}
}

func (uc *UserController) Register(c *gin.Context) {
	var param *model.UserRegisterParam

	err := c.ShouldBindJSON(&param)
	if err != nil {
		gintool.ResError(c, e.INVALID_PARAMS, err)
		return
	}

	err = uc.us.Register(param)
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

	user, err := uc.us.Login(param)
	if err != nil {
		gintool.ResError(c, e.Fail, err)
		return
	}

	users, err := uc.us.GetAllUsersExcept(user.Uid)
	if err != nil {
		gintool.ResError(c, e.Fail, err)
	}

	gintool.ResSuccess(c, map[string]interface{}{
		"loginUser": user,
		"userList":  users,
	})
}
