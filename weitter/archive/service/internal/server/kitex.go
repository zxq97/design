package server

import (
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/archive/kitex_gen/archive/archiveservice"
	"github.com/zxq97/design/weitter/archive/service/internal/service"
)

var ProviderSet = wire.NewSet(NewKiteXServer)

func NewKiteXServer(s *service.ArchiveServiceImpl) server.Server {
	addr, err := net.ResolveTCPAddr("tcp", ":11000")
	if err != nil {
		panic(err)
	}

	return archiveservice.NewServer(s,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "archive"}),
		server.WithServiceAddr(addr),
	)
}
