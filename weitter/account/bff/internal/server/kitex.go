package server

import (
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/account/bff/internal/service"
	"github.com/zxq97/design/weitter/kitex_gen/account/accountbff"
)

var ProviderSet = wire.NewSet(NewKiteX)

func NewKiteX(s *service.AccountBFFImpl) server.Server {
	addr, err := net.ResolveTCPAddr("tcp", ":12009")
	if err != nil {
		panic(err)
	}

	return accountbff.NewServer(s,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "account_bff"}),
		server.WithServiceAddr(addr),
	)
}
