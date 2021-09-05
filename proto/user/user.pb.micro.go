// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user/user.proto

package user

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"

	context "context"

	client "github.com/micro/go-micro/v2/client"

	server "github.com/micro/go-micro/v2/server"
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
	// 注册
	Register(ctx context.Context, in *UserRegisterReq, opts ...client.CallOption) (*UserRegisterResp, error)
	//登录
	Login(ctx context.Context, in *UserLoginReq, opts ...client.CallOption) (*UserLoginResp, error)
	// 查询用户信息
	GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...client.CallOption) (*UserInfoResp, error)
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
		name = "go.micro.server.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Register(ctx context.Context, in *UserRegisterReq, opts ...client.CallOption) (*UserRegisterResp, error) {
	req := c.c.NewRequest(c.name, "User.Register", in)
	out := new(UserRegisterResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *UserLoginReq, opts ...client.CallOption) (*UserLoginResp, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(UserLoginResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...client.CallOption) (*UserInfoResp, error) {
	req := c.c.NewRequest(c.name, "User.GetUserInfo", in)
	out := new(UserInfoResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	// 注册
	Register(context.Context, *UserRegisterReq, *UserRegisterResp) error
	//登录
	Login(context.Context, *UserLoginReq, *UserLoginResp) error
	// 查询用户信息
	GetUserInfo(context.Context, *UserInfoReq, *UserInfoResp) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Register(ctx context.Context, in *UserRegisterReq, out *UserRegisterResp) error
		Login(ctx context.Context, in *UserLoginReq, out *UserLoginResp) error
		GetUserInfo(ctx context.Context, in *UserInfoReq, out *UserInfoResp) error
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

func (h *userHandler) Register(ctx context.Context, in *UserRegisterReq, out *UserRegisterResp) error {
	return h.UserHandler.Register(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *UserLoginReq, out *UserLoginResp) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) GetUserInfo(ctx context.Context, in *UserInfoReq, out *UserInfoResp) error {
	return h.UserHandler.GetUserInfo(ctx, in, out)
}