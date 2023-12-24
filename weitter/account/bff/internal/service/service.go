package service

import (
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/account/bff/internal/biz"
)

var ProviderSet = wire.NewSet(NewAccountBFF)

// AccountBFFImpl implements the last service interface defined in the IDL.
type AccountBFFImpl struct {
	au *biz.AccountBFFUseCase
}

func NewAccountBFF(au *biz.AccountBFFUseCase) *AccountBFFImpl {
	return &AccountBFFImpl{au: au}
}
