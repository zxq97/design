package server

import (
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/like/kitex_gen/like/likeservice"
	"github.com/zxq97/design/weitter/like/service/internal/service"
)

var ProviderSet = wire.NewSet(NewKiteXServer)

func NewKiteXServer(s *service.LikeServiceImpl) server.Server {
	addr, err := net.ResolveTCPAddr("tcp", ":13000")
	if err != nil {
		panic(err)
	}

	return likeservice.NewServer(s,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "like"}),
		server.WithServiceAddr(addr),
	)
}
