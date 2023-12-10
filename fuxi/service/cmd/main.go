package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/zxq97/design/fuxi/service/internal/biz"
)

func newApp(uc *biz.FuxiUseCase, gs *grpc.Server) *kratos.App {
	if err := uc.Load(false); err != nil {
		panic(err)
	}
	return kratos.New(kratos.Server(gs))
}

func main() {
	app := initApp()
	if err := app.Run(); err != nil {
		_ = app.Stop()
	}
}
