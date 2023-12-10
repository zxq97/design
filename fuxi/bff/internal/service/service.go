package service

import (
	"github.com/google/wire"
	"github.com/zxq97/design/fuxi/api/bff/v1"
	"github.com/zxq97/design/fuxi/bff/internal/biz"
)

var ProviderSet = wire.NewSet(NewFuxiBFFService)

type FuxiBFFService struct {
	v1.UnimplementedFuxiBFFServer
	fu *biz.FuxiBFFUseCase
}

func NewFuxiBFFService(fu *biz.FuxiBFFUseCase) *FuxiBFFService {
	return &FuxiBFFService{fu: fu}
}
