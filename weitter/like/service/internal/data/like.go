package data

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/zxq97/design/weitter/like/pkg/constant"
	"github.com/zxq97/design/weitter/like/pkg/model"
	"github.com/zxq97/design/weitter/like/pkg/query"
	"github.com/zxq97/design/weitter/like/service/internal/biz"
	"github.com/zxq97/gokit/pkg/cast"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

var _ biz.LikeRepo = (*likeRepo)(nil)

type likeRepo struct {
	redis *redis.Client
	q     *query.Query
	p     sarama.SyncProducer
}

func NewLikeRepo(redis *redis.Client, db *gorm.DB, p sarama.SyncProducer) biz.LikeRepo {
	return &likeRepo{redis: redis, q: query.Use(db), p: p}
}

func (r *likeRepo) Like(ctx context.Context, objID, uid, author int64, objType int32) error {
	// FIXME kafka msg
	obj := &model.Like{
		ObjID:   objID,
		ObjType: objType,
		UID:     uid,
		Author:  author,
	}
	bs, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, _, err = r.p.SendMessage(&sarama.ProducerMessage{
		Topic: constant.TopicLikeActive,
		Key:   sarama.ByteEncoder(cast.FormatInt(objID)),
		Value: sarama.ByteEncoder(bs),
	})
	return err
}

func (r *likeRepo) Unlike(ctx context.Context, objID, uid int64, objType int32) error {
	// TODO
	return nil
}

func (r *likeRepo) GetLikedUsers(ctx context.Context, objID, lastUID int64, objType int32, limit int8, preload bool) ([]int64, bool, error) {
	uids, err := r.getLikedUsersFromCache(ctx, objID, lastUID, objType, limit+1)
	var hasMore = true
	if err != nil {
		var ls []*model.Like
		ls, hasMore, err = r.getLikedUsersFromDB(ctx, objID, lastUID, objType, limit+1)
		if err != nil {
			return nil, false, err
		}
		r.setLikedUsers(ctx, objID, objType, ls)
		if preload && hasMore {
			// TODO send msg rebuild cache
		}
		uids = make([]int64, len(ls))
		for i := range ls {
			uids[i] = ls[i].UID
		}
	}
	return uids, hasMore, nil
}

func (r *likeRepo) MGetLikedState(ctx context.Context, m map[int32][]int64, uid int64) (map[int32]map[int64]bool, error) {
	var lock sync.Mutex
	res := make(map[int32]map[int64]bool)
	eg, _ := errgroup.WithContext(ctx)
	for k, v := range m {
		k, v := k, v
		eg.Go(func() error {
			ls, missed, err := r.mGetLikedStateFromCache(ctx, v, k, uid)
			if err != nil || len(missed) != 0 {
				dbls, err := r.mGetLikedStateFromDB(ctx, missed, k, uid)
				if err != nil {
					return err
				}
				for i, s := range dbls {
					ls[i] = s
				}
				r.mSetLikedState(ctx, dbls, k, uid)
			}
			lock.Lock()
			for i, s := range ls {
				res[k][i] = s
			}
			defer lock.Unlock()
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *likeRepo) MGetLikedCount(ctx context.Context, m map[int32][]int64) (map[int32]map[int64]int32, error) {
	var lock sync.Mutex
	res := make(map[int32]map[int64]int32)
	eg, _ := errgroup.WithContext(ctx)
	for k, v := range m {
		k, v := k, v
		eg.Go(func() error {
			cm, missed, err := r.mGetLikedCountFromCache(ctx, v, k)
			if err != nil || len(missed) != 0 {
				dbm, err := r.mGetLikedCountFromDB(ctx, missed, k)
				if err != nil {
					return err
				}
				for i, c := range dbm {
					cm[i] = c
				}
				r.mSetLikedCount(ctx, dbm, k)
			}
			lock.Lock()
			for i, c := range cm {
				res[k][i] = c
			}
			defer lock.Unlock()
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return res, nil
}

func (r *likeRepo) MUpdateLikesCount(ctx context.Context, m map[int32]map[int64]int32) error {
	// TODO
	return nil
}

func (r *likeRepo) getLikedUsersFromDB(ctx context.Context, objID, lastUID int64, objType int32, limit int8) (ls []*model.Like, hasMore bool, err error) {
	if lastUID == 0 {
		ls, err = r.q.WithContext(ctx).Like.Select(r.q.Like.UID).Where(r.q.Like.ObjID.Eq(objID), r.q.Like.ObjType.Eq(objType), r.q.Like.State.Eq(1)).Order(r.q.Like.ID.Desc()).Limit(int(limit)).Find()
	} else {
		ls, err = r.q.WithContext(ctx).Like.FindLikedUsers(objID, lastUID, objType, limit)
	}
	hasMore = len(ls) == int(limit)
	return
}

func (r *likeRepo) getLikedUsersFromCache(ctx context.Context, objID, lastUID int64, objType int32, limit int8) ([]int64, error) {
	// TODO need lua
	// XXX revrank lastuid zrevrange(xx, xx+limit)
	return nil, nil
}

func (r *likeRepo) setLikedUsers(ctx context.Context, objID int64, objType int32, ls []*model.Like) {
	key := fmt.Sprintf("lku:%d:%d", objID, objType)
	if r.redis.Expire(ctx, key, time.Hour*4).Val() {
		zs := make([]*redis.Z, len(ls))
		for i := range ls {
			zs[i] = &redis.Z{Member: cast.FormatInt(ls[i].UID), Score: float64(ls[i].Ctime.UnixMilli())}
		}
		_ = r.redis.ZAdd(ctx, key, zs...).Err()
	}
}

func (r *likeRepo) mGetLikedStateFromDB(ctx context.Context, objIDs []int64, objType int32, uid int64) (map[int64]bool, error) {
	ls, err := r.q.WithContext(ctx).Like.Select(r.q.Like.ObjID).Where(r.q.Like.ObjID.In(objIDs...), r.q.Like.ObjType.Eq(objType), r.q.Like.UID.Eq(uid), r.q.Like.State.Eq(1)).Find()
	if err != nil {
		return nil, err
	}
	m := make(map[int64]bool, len(objIDs))
	for i := range ls {
		m[ls[i].ObjID] = true
	}
	for i := range objIDs {
		if _, ok := m[objIDs[i]]; !ok {
			m[objIDs[i]] = false
		}
	}
	return m, nil
}

func (r *likeRepo) mGetLikedStateFromCache(ctx context.Context, objIDs []int64, objType int32, uid int64) (map[int64]bool, []int64, error) {
	keys := make([]string, len(objIDs))
	for i := range objIDs {
		keys[i] = fmt.Sprintf("lkp:%d:%d:%d", objIDs[i], objType, uid)
	}
	res, err := r.redis.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, objIDs, err
	}
	m := make(map[int64]bool, len(objIDs))
	missed := make([]int64, 0, len(objIDs))
	for i := range res {
		if res[i] == nil {
			missed = append(missed, objIDs[i])
			continue
		}
		if val, ok := res[i].(string); !ok {
			missed = append(missed, objIDs[i])
		} else {
			if cast.Atoi(val, 0) > 0 {
				m[objIDs[i]] = true
			} else {
				m[objIDs[i]] = false
			}
		}
	}
	return m, missed, nil
}

func (r *likeRepo) mSetLikedState(ctx context.Context, m map[int64]bool, objType int32, uid int64) {
	pipe := r.redis.Pipeline()
	for k, v := range m {
		ttl := time.Hour * 8
		val := 1
		if !v {
			ttl = time.Hour
			val = 0
		}
		pipe.Set(ctx, fmt.Sprintf("lkp:%d:%d:%d", k, objType, uid), val, ttl)
	}
	_, _ = pipe.Exec(ctx)
}

func (r *likeRepo) mGetLikedCountFromDB(ctx context.Context, objIDs []int64, objType int32) (map[int64]int32, error) {
	cs, err := r.q.WithContext(ctx).LikeCount.Select(r.q.LikeCount.ObjID, r.q.LikeCount.Count).Where(r.q.LikeCount.ObjID.In(objIDs...), r.q.LikeCount.ObjType.Eq(objType)).Find()
	if err != nil {
		return nil, err
	}
	m := make(map[int64]int32, len(objIDs))
	for i := range cs {
		m[cs[i].ObjID] = cs[i].Count
	}
	return m, nil
}

func (r *likeRepo) mGetLikedCountFromCache(ctx context.Context, objIDs []int64, objType int32) (map[int64]int32, []int64, error) {
	keys := make([]string, len(objIDs))
	for i := range objIDs {
		keys[i] = fmt.Sprintf("lkc:%d:%d", objType, objIDs[i])
	}
	res, err := r.redis.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, objIDs, err
	}
	m := make(map[int64]int32, len(objIDs))
	missed := make([]int64, 0, len(objIDs))
	for i := range res {
		if res[i] == nil {
			missed = append(missed, objIDs[i])
			continue
		}
		if val, ok := res[i].(string); !ok {
			missed = append(missed, objIDs[i])
			continue
		} else {
			m[objIDs[i]] = cast.Atoi(val, 0)
		}
	}
	return m, missed, nil
}

func (r *likeRepo) mSetLikedCount(ctx context.Context, m map[int64]int32, objType int32) {
	pipe := r.redis.Pipeline()
	for k, v := range m {
		pipe.Set(ctx, fmt.Sprintf("lkc:%d:%d", objType, k), v, time.Hour*8)
	}
	_, _ = pipe.Exec(ctx)
}
