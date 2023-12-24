package data

import (
	"context"

	"github.com/zxq97/design/weitter/account/bff/internal/biz"
	"github.com/zxq97/design/weitter/account/kitex_gen/account"
	"github.com/zxq97/design/weitter/account/kitex_gen/account/accountservice"
)

var _ biz.AccountBFFRepo = (*accountBFFRepo)(nil)

type accountBFFRepo struct {
	client accountservice.Client
}

func NewAccountBFFRepo(client accountservice.Client) biz.AccountBFFRepo {
	return &accountBFFRepo{client: client}
}

func (r *accountBFFRepo) CreateUser(ctx context.Context, user *biz.User) error {
	_, err := r.client.CreateUser(ctx, &account.CreateUserRequest{
		User: &account.User{
			Uid:          user.UID,
			Nickname:     user.Nickname,
			Gender:       account.Gender(user.Gender),
			Introduction: user.Introduction,
		},
	})
	return err
}

func (r *accountBFFRepo) MGetUsers(ctx context.Context, uids []int64) (map[int64]*biz.User, error) {
	res, err := r.client.MGetUsers(ctx, &account.MGetUsersRequest{
		Uids: uids,
	})
	if err != nil {
		return nil, err
	}
	m := make(map[int64]*biz.User, len(res.Users))
	for k, v := range res.Users {
		m[k] = &biz.User{
			UID:          v.Uid,
			Nickname:     v.Nickname,
			Gender:       int32(v.Gender),
			Introduction: v.Introduction,
		}
	}
	return m, nil
}

func (r *accountBFFRepo) DeleteUsers(ctx context.Context, uids []int64) error {
	_, err := r.client.DeleteUsers(ctx, &account.DeleteUsersRequest{
		Uids: uids,
	})
	return err
}
