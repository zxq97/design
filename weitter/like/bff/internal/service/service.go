package service

import (
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/like/bff/internal/biz"
)

var ProviderSet = wire.NewSet(NewLikeBFF)

// LikeBFFImpl implements the last service interface defined in the IDL.
type LikeBFFImpl struct {
	lu *biz.LikeBFFUseCase
}

func NewLikeBFF(lu *biz.LikeBFFUseCase) *LikeBFFImpl {
	return &LikeBFFImpl{lu: lu}
}
