package core

import (
	"context"

	"myproject/data"
	"myproject/data/dbset"
	"myproject/internal/mapping"
	"myproject/internal/micro/client"
	"myproject/internal/model/entity"
	"myproject/internal/proto/basic"
	"myproject/internal/proto/hello"
	bsiErrors "bsi/kratos/micro/errors"
	mgorm "bsi/kratos/micro/gorm"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type UserService struct {
	*mgorm.TransactionScope
	client *client.Client

	repo *data.UserRepo
	log  *log.Helper
}

func NewUserService(client *client.Client, repo *data.UserRepo, logger log.Logger, scope *mgorm.TransactionScope) *UserService {
	return &UserService{client: client, repo: repo, log: log.NewHelper(logger), TransactionScope: scope}
}

func (s *UserService) FindByID(ctx context.Context, in *basic.Int64) (*hello.UserInfo, error) {
	user, err := s.repo.FindUserByID(ctx, in.Value)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, bsiErrors.MapError(ctx, int32(hello.Errors_Errors_NotFoundError_Basic))
		}
		return nil, err
	}
	info := &hello.UserInfo{}
	if err = mapping.Copy(info, user); err != nil {
		return nil, err
	}
	return info, nil
}

func (s *UserService) Update(ctx context.Context, in *hello.UserInfo) (*hello.UserInfo, error) {
	updates := map[string]any{
		"id":     in.Id,
		"status": in.Status,
	}
	user, err := s.repo.Update(ctx, updates)
	if err != nil {
		return nil, err
	}
	info := &hello.UserInfo{}
	if err = mapping.Copy(info, user); err != nil {
		return nil, err
	}
	return info, nil
}

func (s *UserService) List(ctx context.Context, in *basic.PageRequest) (*hello.UserListResponse, error) {
	total, users, err := s.repo.List(ctx, int(in.Index), int(in.Size))
	if err != nil {
		return nil, err
	}
	info := &hello.UserListResponse{}
	info.Total = total
	info.Items = make([]*hello.UserInfo, 0, len(users))
	for _, user := range users {
		info.Items = append(info.Items, &hello.UserInfo{
			Id:     user.ID,
			Status: hello.UserStatus(user.Status),
		})
	}
	return info, nil
}

// Notice: Core 有事务的 Example。同时 Data 层也可以有事务。
func (s *UserService) TransactionExample(ctx context.Context, in *hello.UserInfo) error {
	var err error
	user := &dbset.User{
		User: entity.User{
			Name:   in.Name,
			Status: int32(in.Status),
		},
	}
	// 其他业务操作
	// Some business logic code

	// 开启事务
	err = s.Transaction(ctx, func(ctx context.Context) error {
		user.ID = 1
		err = s.repo.Create(ctx, user)
		if err != nil {
			return err
		}

		// 这里模拟其他 Repo 的更新操作
		_, err = s.repo.Update(ctx, map[string]any{"id": user.ID, "status": user.Status})
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
