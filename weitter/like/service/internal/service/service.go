package service

import (
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/like/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewLikeService)

// LikeServiceImpl implements the last service interface defined in the IDL.
type LikeServiceImpl struct {
	lu *biz.LikeUseCase
}

func NewLikeService(lu *biz.LikeUseCase) *LikeServiceImpl {
	return &LikeServiceImpl{lu: lu}
}
