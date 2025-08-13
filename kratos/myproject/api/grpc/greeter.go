package grpc

import (
	"myproject/core"
	"myproject/internal/proto/basic"
	"myproject/internal/proto/hello"
	"bsi/kratos/micro/errors"
	"bsi/kratos/micro/linker"
	"context"
)

type GreeterHandler struct {
	hello.UnimplementedGreeterServer

	greeterSrv *core.GreeterService
}

func NewGreeterHandler(greeterSrv *core.GreeterService) *GreeterHandler {
	return &GreeterHandler{greeterSrv: greeterSrv}
}

// @anonymous
func (s *GreeterHandler) SayHello(ctx context.Context, in *basic.String) (*basic.String, error) {
	if in.Value == "throw" {
		//抛一个普通error => 中间件转为micro/errors/Error
		//return nil, errors.New("test error")

		//直接抛一个micro/errors/Error
		//return nil, errors.New("测试错误消息格式")

		//跟进错误码自动获取错误消息(多语言)
		return nil, errors.MapError(ctx, int32(hello.Errors_Errors_AuthorizationError_Basic))
	}
	return s.greeterSrv.SayHello(ctx, in)
}

// @anonymous
func (s *GreeterHandler) TestLink(ctx context.Context, in *hello.TestLinkDto) (*hello.TestLinkDto, error) {
	//link单个对象
	err := linker.Link(ctx, in)

	//link集合对象
	//err = linker.Link(ctx, []*hello.TestLinkDto{in, in})

	return in, err
}
