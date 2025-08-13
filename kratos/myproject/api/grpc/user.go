package grpc

import (
	"context"

	"myproject/core"
	"myproject/internal/proto/basic"
	"myproject/internal/proto/hello"
)

type UserHandler struct {
	hello.UnimplementedUserServer

	userSrv *core.UserService
}

func NewUserHandler(userSrv *core.UserService) *UserHandler {
	return &UserHandler{userSrv: userSrv}
}

func (s *UserHandler) FindByID(ctx context.Context, in *basic.Int64) (*hello.UserInfo, error) {
	return s.userSrv.FindByID(ctx, in)
}

// FindList
// @name 用户.管理.查询列表
// @desc 查询用户列表
func (s *UserHandler) FindList(ctx context.Context, in *basic.PageRequest) (*hello.UserListResponse, error) {
	return s.userSrv.List(ctx, in)
}

// Update
// @name 用户.管理.更新
// @desc 更新用户信息
func (s *UserHandler) Update(ctx context.Context, in *hello.UserInfo) (*hello.UserInfo, error) {
	return s.userSrv.Update(ctx, in)
}
