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
	//say hello
	Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 新增用户
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*CreateUserResponse, error)
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
		name = "com.txvier.srv.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "User.Hello", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...client.CallOption) (*CreateUserResponse, error) {
	req := c.c.NewRequest(c.name, "User.CreateUser", in)
	out := new(CreateUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	//say hello
	Hello(context.Context, *Request, *Response) error
	// 新增用户
	CreateUser(context.Context, *CreateUserRequest, *CreateUserResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Hello(ctx context.Context, in *Request, out *Response) error
		CreateUser(ctx context.Context, in *CreateUserRequest, out *CreateUserResponse) error
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

func (h *userHandler) Hello(ctx context.Context, in *Request, out *Response) error {
	return h.UserHandler.Hello(ctx, in, out)
}

func (h *userHandler) CreateUser(ctx context.Context, in *CreateUserRequest, out *CreateUserResponse) error {
	return h.UserHandler.CreateUser(ctx, in, out)
}
