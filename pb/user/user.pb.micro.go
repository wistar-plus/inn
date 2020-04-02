// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package user

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

// Client API for User service

type UserService interface {
	//rpc GetUserById(GetUserByIdRequest) returns (UserResponse){}
	GetAllUsersExcept(ctx context.Context, in *GetAllUsersExceptRequest, opts ...client.CallOption) (*UsersResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*UserResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) GetAllUsersExcept(ctx context.Context, in *GetAllUsersExceptRequest, opts ...client.CallOption) (*UsersResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetAllUsersExcept", in)
	out := new(UsersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.name, "User.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	//rpc GetUserById(GetUserByIdRequest) returns (UserResponse){}
	GetAllUsersExcept(context.Context, *GetAllUsersExceptRequest, *UsersResponse) error
	Login(context.Context, *LoginRequest, *UserResponse) error
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		GetAllUsersExcept(ctx context.Context, in *GetAllUsersExceptRequest, out *UsersResponse) error
		Login(ctx context.Context, in *LoginRequest, out *UserResponse) error
		Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) GetAllUsersExcept(ctx context.Context, in *GetAllUsersExceptRequest, out *UsersResponse) error {
	return h.UserHandler.GetAllUsersExcept(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *LoginRequest, out *UserResponse) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.UserHandler.Register(ctx, in, out)
}