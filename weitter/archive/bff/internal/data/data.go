package data

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/cloudwego/kitex/client"
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/archive/kitex_gen/archive/archiveservice"
	"github.com/zxq97/design/weitter/kitex_gen/account/accountbff"
)

var ProviderSet = wire.NewSet(NewMemCache, NewArchiveClient, NewAccountClient, NewArchiveBFFRepo)

func NewMemCache() *memcache.Client {
	return memcache.New([]string{"127.0.0.1:11211"}...)
}

func NewArchiveClient() archiveservice.Client {
	cli, err := archiveservice.NewClient("archive",
		client.WithHostPorts(":11000"),
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
