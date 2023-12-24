package data

import (
	"github.com/cloudwego/kitex/client"
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/account/kitex_gen/account/accountservice"
)

var ProviderSet = wire.NewSet(NewAccountClient, NewAccountBFFRepo)

func NewAccountClient() accountservice.Client {
	cli, err := accountservice.NewClient("account",
		client.WithHostPorts(":12000"),
	)
	if err != nil {
		panic(err)
	}
	return cli
}
