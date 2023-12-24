package service

import (
	"github.com/google/wire"
	"github.com/zxq97/design/weitter/archive/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewArchiveService)

// ArchiveServiceImpl implements the last service interface defined in the IDL.
type ArchiveServiceImpl struct {
	au *biz.ArchiveUseCase
}

func NewArchiveService(au *biz.ArchiveUseCase) *ArchiveServiceImpl {
	return &ArchiveServiceImpl{au: au}
}
