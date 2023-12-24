package server

import (
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/account/kitex_gen/account/accountservice"
	"github.com/zxq97/design/weitter/account/service/internal/service"
)

var ProviderSet = wire.NewSet(NewKiteX)

func NewKiteX(s *service.AccountServiceImpl) server.Server {
	addr, err := net.ResolveTCPAddr("tcp", ":12000")
	if err != nil {
		panic(err)
	}

	return accountservice.NewServer(s,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "account"}),
		server.WithServiceAddr(addr),
	)
}
