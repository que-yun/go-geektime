package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "go-geektime/fourth-week/api/helloworld/v1"
	"go-geektime/fourth-week/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, errors.New("not fount ")
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
