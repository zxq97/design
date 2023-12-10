package server

import (
	"time"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/zxq97/design/fuxi/api/bff/v1"
	"github.com/zxq97/design/fuxi/bff/internal/service"
)

func NewHttpServer(s *service.FuxiBFFService) *http.Server {
	var opts = []http.ServerOption{
		http.Address(":8000"),
		http.Middleware(recovery.Recovery()),
		http.Timeout(time.Second),
	}

	svr := http.NewServer(opts...)
	v1.RegisterFuxiBFFHTTPServer(svr, s)
	return svr
}
