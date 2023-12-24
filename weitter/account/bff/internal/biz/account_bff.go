package biz

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zxq97/design/weitter/account/kitex_gen/account"
)

type User struct {
	UID          int64
	Nickname     string
	Gender       int32
	Introduction string
}

func (uc *AccountBFFUseCase) CreateUser(ctx context.Context, user *User) error {
	user.UID = uc.idGen.Gen()
	return uc.repo.CreateUser(ctx, user)
}

func (uc *AccountBFFUseCase) GetUser(ctx context.Context, uid int64) (*User, error) {
	m, err := uc.repo.MGetUsers(ctx, []int64{uid})
	if err != nil {
		return nil, err
	}
	u, ok := m[uid]
	if !ok {
		return nil, errors.New("account_bff: user not found")
	}
	return u, nil
}

func (uc *AccountBFFUseCase) MGetUsers(ctx context.Context, uids []int64) (map[int64]*User, error) {
	return uc.repo.MGetUsers(ctx, uids)
}

func (uc *AccountBFFUseCase) DeleteUsers(ctx context.Context, uid int64, uids []int64) error {
	if len(uids) > 1 {
		u, err := uc.GetUser(ctx, uid)
		if err != nil {
			return err
		}
		if u.Gender != int32(account.Gender_Admin) {
			return errors.New("account_bff: no permission delete")
		}
	} else if uid != uids[0] {
		return errors.New("account_bff: no permission delete")
	}
	return uc.repo.DeleteUsers(ctx, uids)
}
