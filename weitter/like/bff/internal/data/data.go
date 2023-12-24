package data

import (
	"github.com/cloudwego/kitex/client"
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/kitex_gen/account/accountbff"
	"github.com/zxq97/design/weitter/like/kitex_gen/like/likeservice"
)

var ProviderSet = wire.NewSet(NewLikeClient, NewAccountClient, NewLikeBFFRepo)

func NewLikeClient() likeservice.Client {
	cli, err := likeservice.NewClient("like",
		client.WithHostPorts(":13000"),
	)
	if err != nil {
		panic(err)
	}
	return cli
}

func NewAccountClient() accountbff.Client {
	cli, err := accountbff.NewClient("account_bff",
		client.WithHostPorts("12009"),
	)
	if err != nil {
		panic(err)
	}
	return cli
}
