// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/zxq97/design/fuxi/service/internal/biz"
	"github.com/zxq97/design/fuxi/service/internal/data"
	"github.com/zxq97/design/fuxi/service/internal/server"
	"github.com/zxq97/design/fuxi/service/internal/service"
)

// Injectors from wire.go:

func initApp() *kratos.App {
	db := data.NewDB()
	fuxiRepo := data.NewFuxiRepo(db)
	fuxiUseCase := biz.NewFuxiUseCase(fuxiRepo)
	fuxiService := service.NewFuxiService(fuxiUseCase)
	grpcServer := server.NewGrpcServer(fuxiService)
	app := newApp(fuxiUseCase, grpcServer)
	return app
}