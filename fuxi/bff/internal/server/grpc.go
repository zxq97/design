package server

import (
	"time"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/zxq97/design/fuxi/api/bff/v1"
	"github.com/zxq97/design/fuxi/bff/internal/service"
)

var ProviderSet = wire.NewSet(NewGrpcServer, NewHttpServer)

func NewGrpcServer(s *service.FuxiBFFService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Address(":8888"),
		grpc.Middleware(recovery.Recovery()),
		grpc.Timeout(time.Second),
	}

	svr := grpc.NewServer(opts...)
	v1.RegisterFuxiBFFServer(svr, s)
	return svr
}
