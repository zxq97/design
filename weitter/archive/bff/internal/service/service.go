package service

import (
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/archive/bff/internal/biz"
)

var ProviderSet = wire.NewSet(NewArchiveBFF)

// ArchiveBFFImpl implements the last service interface defined in the IDL.
type ArchiveBFFImpl struct {
	au *biz.ArchiveBFFUseCase
}

func NewArchiveBFF(au *biz.ArchiveBFFUseCase) *ArchiveBFFImpl {
	return &ArchiveBFFImpl{au: au}
}
