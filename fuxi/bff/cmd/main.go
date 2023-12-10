package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func newApp(gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(kratos.Server(gs, hs))
}

func main() {
	app := initApp()
	if err := app.Run(); err != nil {
		_ = app.Stop()
	}
}
