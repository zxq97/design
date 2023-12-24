package service

import (
	"context"

	"github.com/zxq97/design/weitter/account/kitex_gen/account"
	"github.com/zxq97/design/weitter/account/pkg/model"
)

// CreateUser implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) CreateUser(ctx context.Context, req *account.CreateUserRequest) (resp *account.EmptyResponse, err error) {
	resp = new(account.EmptyResponse)
	err = s.au.CreateUser(ctx, &model.User{
		UID:          req.User.Uid,
		Nickname:     req.User.Nickname,
		Gender:       int32(req.User.Gender),
		Introduction: req.User.Introduction,
	})
	return
}

// MGetUsers implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) MGetUsers(ctx context.Context, req *account.MGetUsersRequest) (resp *account.MGetUsersResponse, err error) {
	resp = new(account.MGetUsersResponse)
	m, err := s.au.MGetUsers(ctx, req.Uids)
	if err != nil {
		return resp, err
	}
	resp.Users = make(map[int64]*account.User, len(m))
	for k, v := range m {
		resp.Users[k] = &account.User{
			Uid:          v.UID,
			Nickname:     v.Nickname,
			Gender:       account.Gender(v.Gender),
			Introduction: v.Introduction,
		}
	}
	return
}

// DeleteUsers implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) DeleteUsers(ctx context.Context, req *account.DeleteUsersRequest) (resp *account.EmptyResponse, err error) {
	resp = new(account.EmptyResponse)
	err = s.au.DeleteUsers(ctx, req.Uids)
	return
}
