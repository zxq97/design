package data

import (
	"context"
	"encoding/json"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/zxq97/design/weitter/account/pkg/model"
	"github.com/zxq97/design/weitter/account/pkg/query"
	"github.com/zxq97/design/weitter/account/service/internal/biz"
	"github.com/zxq97/gokit/pkg/cast"
	"gorm.io/gorm"
)

var _ biz.AccountRepo = (*accountRepo)(nil)

type accountRepo struct {
	mc *memcache.Client
	q  *query.Query
}

func NewAccountRepo(mc *memcache.Client, db *gorm.DB) biz.AccountRepo {
	return &accountRepo{mc: mc, q: query.Use(db)}
}

func (r *accountRepo) CreateUser(ctx context.Context, user *model.User) error {
	return r.q.WithContext(ctx).User.Create(user)
}

func (r *accountRepo) MGetUsers(ctx context.Context, uids []int64) (map[int64]*model.User, error) {
	m, missed, err := r.mGetUsersFromCache(ctx, uids)
	if err != nil || len(missed) != 0 {
		dbm, err := r.mGetUsersFromDB(ctx, missed)
		if err != nil {
			return nil, err
		}
		for k, v := range dbm {
			m[k] = v
		}
		r.mSetUsers(ctx, dbm)
	}
	return m, nil
}

func (r *accountRepo) DeleteUsers(ctx context.Context, uids []int64) error {
	if _, err := r.q.WithContext(ctx).User.Where(r.q.User.UID.In(uids...), r.q.User.Status.Eq(0)).Update(r.q.User.Status, 1); err != nil {
		return err
	}
	r.deleteUsers(ctx, uids)
	return nil
}

func (r *accountRepo) mGetUsersFromCache(ctx context.Context, uids []int64) (map[int64]*model.User, []int64, error) {
	keys := make([]string, len(uids))
	for i := range uids {
		keys[i] = "aco:" + cast.FormatInt(uids[i])
	}
	res, err := r.mc.GetMulti(keys)
	if err != nil {
		return nil, uids, err
	}
	m := make(map[int64]*model.User, len(uids))
	missed := make([]int64, 0, len(uids))
	for _, v := range res {
		var a model.User
		if err = json.Unmarshal(v.Value, &a); err != nil {
			continue
		}
		m[a.UID] = &a
	}
	for i := range uids {
		if _, ok := m[uids[i]]; !ok {
			missed = append(missed, uids[i])
		}
	}
	return m, missed, nil
}

func (r *accountRepo) mGetUsersFromDB(ctx context.Context, uids []int64) (map[int64]*model.User, error) {
	us, err := r.q.WithContext(ctx).User.Select(r.q.User.UID, r.q.User.Nickname, r.q.User.Gender, r.q.User.Introduction).Where(r.q.User.UID.In(uids...), r.q.User.Status.Eq(0)).Find()
	if err != nil {
		return nil, err
	}
	m := make(map[int64]*model.User, len(uids))
	for i := range us {
		m[us[i].UID] = us[i]
	}
	return m, nil
}

func (r *accountRepo) mSetUsers(ctx context.Context, m map[int64]*model.User) {
	for k, v := range m {
		bs, err := json.Marshal(v)
		if err != nil {
			continue
		}
		_ = r.mc.Set(&memcache.Item{Key: "aco:" + cast.FormatInt(k), Value: bs, Expiration: 8 * 3600})
	}
}

func (r *accountRepo) deleteUsers(ctx context.Context, uids []int64) {
	for i := range uids {
		_ = r.mc.Delete("aco:" + cast.FormatInt(uids[i]))
	}
}
