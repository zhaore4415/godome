package core

import (
	"context"

	"myproject/data"
	"myproject/internal/micro/client"
	"myproject/internal/proto/basic"

	"github.com/go-kratos/kratos/v2/log"
)

type GreeterService struct {
	client *client.Client

	repo *data.UserRepo
	log  *log.Helper
}

func NewGreeterService(client *client.Client, repo *data.UserRepo, logger log.Logger) *GreeterService {
	return &GreeterService{client: client, repo: repo, log: log.NewHelper(logger)}
}

func (s *GreeterService) SayHello(ctx context.Context, in *basic.String) (*basic.String, error) {
	s.log.WithContext(ctx).Infof("SayHello: %v", in.Value)
	return &basic.String{Value: in.Value}, nil
}
