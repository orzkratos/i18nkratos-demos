package service

import (
	"context"

	v1 "github.com/orzkratos/demokratos/demo1kratos/api/helloworld/v1"
	"github.com/orzkratos/demokratos/demo1kratos/internal/biz"
	"github.com/orzkratos/demokratos/demo1kratos/internal/pkg/middleware/localize"
	"github.com/orzkratos/demokratos/demo1kratos/internal/pkg/middleware/localize/i18nmsg"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}

	message := localize.FromContext(ctx).MustLocalize(i18nmsg.I18nGreeting(&i18nmsg.GreetingParam{
		Name: g.Hello,
	}))

	return &v1.HelloReply{Message: message}, nil
}
