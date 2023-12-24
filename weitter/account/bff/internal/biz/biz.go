package biz

import (
	"context"

	"github.com/google/wire"
	"github.com/zxq97/gokit/pkg/generate"
)

var ProviderSet = wire.NewSet(NewAccountBFFUseCase)

type AccountBFFRepo interface {
	CreateUser(ctx context.Context, user *User) error
	MGetUsers(ctx context.Context, uids []int64) (map[int64]*User, error)
	DeleteUsers(ctx context.Context, uids []int64) error
}

type AccountBFFUseCase struct {
	repo  AccountBFFRepo
	idGen *generate.SnowIDGen
}

func NewAccountBFFUseCase(repo AccountBFFRepo) *AccountBFFUseCase {
	idGen, err := generate.NewSnowIDGen("account_bff")
	if err != nil {
		panic(err)
	}
	return &AccountBFFUseCase{repo: repo, idGen: idGen}
}
