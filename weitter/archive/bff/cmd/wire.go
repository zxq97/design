//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/archive/bff/internal/biz"
	"github.com/zxq97/design/weitter/archive/bff/internal/data"
	srv "github.com/zxq97/design/weitter/archive/bff/internal/server"
	"github.com/zxq97/design/weitter/archive/bff/internal/service"
)

func initServer() server.Server {
	wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, srv.ProviderSet)
	return nil
}