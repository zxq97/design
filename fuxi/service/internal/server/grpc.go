package server

import (
	"time"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/zxq97/design/fuxi/api/service/v1"
	"github.com/zxq97/design/fuxi/service/internal/service"
)

var ProviderSet = wire.NewSet(NewGrpcServer)

func NewGrpcServer(s *service.FuxiService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Address(":8889"),
		grpc.Middleware(recovery.Recovery()),
		grpc.Timeout(time.Millisecond * 200),
	}

	svr := grpc.NewServer(opts...)
	v1.RegisterFuxiServer(svr, s)
	return svr
}
