package biz

import (
	"context"

	"github.com/google/wire"
	"github.com/zxq97/design/weitter/account/pkg/model"
)

var ProviderSet = wire.NewSet(NewAccountUseCase)

type AccountRepo interface {
	CreateUser(ctx context.Context, user *model.User) error
	MGetUsers(ctx context.Context, uids []int64) (map[int64]*model.User, error)
	DeleteUsers(ctx context.Context, uids []int64) error
}

type AccountUseCase struct {
	repo AccountRepo
}

func NewAccountUseCase(repo AccountRepo) *AccountUseCase {
	return &AccountUseCase{repo: repo}
}
