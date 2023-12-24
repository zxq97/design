package server

import (
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/archive/bff/internal/service"
	"github.com/zxq97/design/weitter/kitex_gen/archive/archivebff"
)

var ProviderSet = wire.NewSet(NewKiteXServer)

func NewKiteXServer(s *service.ArchiveBFFImpl) server.Server {
	addr, err := net.ResolveTCPAddr("tcp", ":11009")
	if err != nil {
		panic(err)
	}

	return archivebff.NewServer(s,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "archive_bff"}),
		server.WithServiceAddr(addr),
	)
}
