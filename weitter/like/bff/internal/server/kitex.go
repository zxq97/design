package server

import (
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/kitex_gen/like/likebff"
	"github.com/zxq97/design/weitter/like/bff/internal/service"
)

var ProviderSet = wire.NewSet(NewKiteX)

func NewKiteX(s *service.LikeBFFImpl) server.Server {
	addr, err := net.ResolveTCPAddr("tcp", ":13009")
	if err != nil {
		panic(err)
	}

	return likebff.NewServer(s,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "like_bff"}),
		server.WithServiceAddr(addr),
	)
}
