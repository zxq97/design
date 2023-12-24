package service

import (
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/account/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewAccountService)

// AccountServiceImpl implements the last service interface defined in the IDL.
type AccountServiceImpl struct {
	au *biz.AccountUseCase
}

func NewAccountService(au *biz.AccountUseCase) *AccountServiceImpl {
	return &AccountServiceImpl{au: au}
}
