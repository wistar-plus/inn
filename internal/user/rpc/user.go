package rpc

import (
	"context"
	"inn/internal/user/model"
	"inn/internal/user/repository"
	userpb "inn/pb/user"
	"inn/pkg/e"
)

type userRpc struct {
	repo repository.IUserRepository
	//ms *service.user
}

func NewUserRpc(repo repository.IUserRepository) *userRpc {
	return &userRpc{repo: repo}
}

func (ur *userRpc) Login(ctx context.Context, req *userpb.LoginRequest, rsp *userpb.UserResponse) error {
	user, err := ur.repo.FindByEmail(req.GetEmail())
	if err != nil {
		return e.ERROR_DBERROR
	}
	if user.Pwd != req.GetPwd() {
		return e.ERROR_WRONGPASSWORD
	}
	rsp.Uid = user.Uid
	rsp.Email = user.Email
	rsp.NickName = user.NickName

	return err
}

func (ur *userRpc) Register(ctx context.Context, req *userpb.RegisterRequest, rsp *userpb.RegisterResponse) error {

	_, err := ur.repo.FindByEmail(req.GetEmail())
	if err != nil {
		return err
	}

	user := &model.User{
		Email:    req.GetEmail(),
		Pwd:      req.GetPwd(),
		NickName: req.GetNickName(),
	}

	return ur.repo.Create(user)
}

func (ur *userRpc) GetAllUsersExcept(ctx context.Context, req *userpb.GetAllUsersExceptRequest, rsp *userpb.UsersResponse) error {
	users, err := ur.repo.FindUsersByUidIsNot(req.GetUid())
	if err != nil {
		return e.ERROR_DBERROR
	}

	rsp.Users = make([]*userpb.UserResponse, len(users))

	for i, user := range users {
		rsp.Users[i] = &userpb.UserResponse{
			Uid:      user.Uid,
			Email:    user.Email,
			NickName: user.NickName,
		}
	}

	return nil
}
