package service

import (
	"context"

	"github.com/zxq97/design/weitter/account/bff/internal/biz"
	"github.com/zxq97/design/weitter/kitex_gen/account"
	"github.com/zxq97/design/weitter/kitex_gen/common"
)

// CreateUser implements the AccountBFFImpl interface.
func (s *AccountBFFImpl) CreateUser(ctx context.Context, req *account.CreateUserRequest) (resp *common.EmptyResponse, err error) {
	resp = new(common.EmptyResponse)
	err = s.au.CreateUser(ctx, &biz.User{
		Nickname:     req.Nickname,
		Gender:       int32(req.Gender),
		Introduction: req.Introduction,
	})
	return
}

// GetUser implements the AccountBFFImpl interface.
func (s *AccountBFFImpl) GetUser(ctx context.Context, req *common.GetItemRequest) (resp *account.GetUserResponse, err error) {
	resp = new(account.GetUserResponse)
	u, err := s.au.GetUser(ctx, req.Id)
	if err != nil {
		return resp, err
	}
	resp.User = &account.User{
		Uid:          u.UID,
		Nickname:     u.Nickname,
		Gender:       account.Gender(u.Gender),
		Introduction: u.Introduction,
	}
	return
}

// MGetUsers implements the AccountBFFImpl interface.
func (s *AccountBFFImpl) MGetUsers(ctx context.Context, req *common.MGetItemARequest) (resp *account.MGetUsersResponse, err error) {
	resp = new(account.MGetUsersResponse)
	m, err := s.au.MGetUsers(ctx, req.Ids)
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

// DeleteUsers implements the AccountBFFImpl interface.
func (s *AccountBFFImpl) DeleteUsers(ctx context.Context, req *account.DeleteUsersRequest) (resp *common.EmptyResponse, err error) {
	resp = new(common.EmptyResponse)
	err = s.au.DeleteUsers(ctx, req.Uid, req.Uids)
	return
}

// CheckAdminUser implements the AccountBFFImpl interface.
func (s *AccountBFFImpl) CheckAdminUser(ctx context.Context, req *common.GetItemRequest) (resp *common.CheckResponse, err error) {
	resp = new(common.CheckResponse)
	u, err := s.au.GetUser(ctx, req.Id)
	if err != nil {
		return resp, err
	}
	resp.Ok = u.Gender == int32(account.Gender_Admin)
	return
}
