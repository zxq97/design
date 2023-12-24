package biz

import (
	"context"

	"github.com/zxq97/design/weitter/account/pkg/model"
)

func (uc *AccountUseCase) CreateUser(ctx context.Context, user *model.User) error {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *AccountUseCase) MGetUsers(ctx context.Context, uids []int64) (map[int64]*model.User, error) {
	return uc.repo.MGetUsers(ctx, uids)
}

func (uc *AccountUseCase) DeleteUsers(ctx context.Context, uids []int64) error {
	return uc.repo.DeleteUsers(ctx, uids)
}
