//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/zxq97/design/fuxi/service/internal/biz"
	"github.com/zxq97/design/fuxi/service/internal/data"
	"github.com/zxq97/design/fuxi/service/internal/server"
	"github.com/zxq97/design/fuxi/service/internal/service"
)

func initApp() *kratos.App {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
