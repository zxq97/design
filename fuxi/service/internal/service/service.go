package service

import (
	"github.com/google/wire"
	"github.com/zxq97/design/fuxi/api/service/v1"
	"github.com/zxq97/design/fuxi/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewFuxiService)

type FuxiService struct {
	v1.UnimplementedFuxiServer
	fu *biz.FuxiUseCase
}

func NewFuxiService(fu *biz.FuxiUseCase) *FuxiService {
	return &FuxiService{fu: fu}
}
