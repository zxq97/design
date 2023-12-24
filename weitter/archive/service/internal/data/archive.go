package data

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis/v8"
	"github.com/zxq97/design/weitter/archive/pkg/constant"
	"github.com/zxq97/design/weitter/archive/pkg/model"
	"github.com/zxq97/design/weitter/archive/pkg/query"
	"github.com/zxq97/design/weitter/archive/service/internal/biz"
	"github.com/zxq97/gokit/pkg/cast"
	"gorm.io/gorm"
)

var _ biz.ArchiveRepo = (*archiveRepo)(nil)

type archiveRepo struct {
	redis *redis.Client
	mc    *memcache.Client
	q     *query.Query
	p     sarama.SyncProducer
}

func NewArchiveRepo(redis *redis.Client, mc *memcache.Client, db *gorm.DB, p sarama.SyncProducer) biz.ArchiveRepo {
	return &archiveRepo{redis: redis, mc: mc, q: query.Use(db), p: p}
}

func (r *archiveRepo) CreateArticle(ctx context.Context, article *model.Article) error {
	bs, err := json.Marshal(article)
	if err != nil {
		return err
	}
	_, _, err = r.p.SendMessage(&sarama.ProducerMessage{
		Topic: constant.TopicCreateArticle,
		Key:   sarama.ByteEncoder(cast.FormatInt(article.UID)),
		Value: sarama.ByteEncoder(bs),
	})
	return err
}

func (r *archiveRepo) MGetArticles(ctx context.Context, articleIDs []int64) (map[int64]*model.Article, error) {
	m, missed, err := r.mGetArticlesFromCache(ctx, articleIDs)
	if err != nil || len(missed) != 0 {
		dbm, err := r.mGetArticlesFromDB(ctx, missed)
		if err != nil {
			return nil, err
		}
		for k, v := range dbm {
			m[k] = v
		}
		r.mSetArticles(ctx, dbm)
	}
	return m, nil
}

func (r *archiveRepo) GetUserArticles(ctx context.Context, uid, lastID int64, limit int8) ([]*model.Article, bool, error) {
	// TODO
	return nil, false, nil
}

func (r *archiveRepo) DeleteArticles(ctx context.Context, articleIDs []int64) error {
	if _, err := r.q.WithContext(ctx).Article.Where(r.q.Article.ArticleID.In(articleIDs...)).Delete(); err != nil {
		return err
	}
	r.deleteArticles(ctx, articleIDs)
	return nil
}

func (r *archiveRepo) mGetArticlesFromCache(ctx context.Context, articleIDs []int64) (map[int64]*model.Article, []int64, error) {
	keys := make([]string, len(articleIDs))
	for i := range articleIDs {
		keys[i] = "art:" + cast.FormatInt(articleIDs[i])
	}
	res, err := r.mc.GetMulti(keys)
	if err != nil {
		return nil, articleIDs, err
	}
	m := make(map[int64]*model.Article, len(articleIDs))
	missed := make([]int64, 0, len(articleIDs))
	for _, v := range res {
		var a model.Article
		if err = json.Unmarshal(v.Value, &a); err != nil {
			continue
		}
		m[a.ArticleID] = &a
	}
	for i := range articleIDs {
		if _, ok := m[articleIDs[i]]; !ok {
			missed = append(missed, articleIDs[i])
		}
	}
	return m, missed, nil
}

func (r *archiveRepo) mGetArticlesFromDB(ctx context.Context, articleIDs []int64) (map[int64]*model.Article, error) {
	as, err := r.q.WithContext(ctx).Article.Where(r.q.Article.ArticleID.In(articleIDs...)).Find()
	if err != nil {
		return nil, err
	}
	m := make(map[int64]*model.Article, len(articleIDs))
	for _, v := range as {
		m[v.ArticleID] = v
	}
	return m, nil
}

func (r *archiveRepo) mSetArticles(ctx context.Context, m map[int64]*model.Article) {
	for k, v := range m {
		bs, err := json.Marshal(v)
		if err != nil {
			continue
		}
		_ = r.mc.Set(&memcache.Item{Key: "art:" + cast.FormatInt(k), Value: bs, Expiration: 8 * 3600})
	}
}

func (r *archiveRepo) deleteArticles(ctx context.Context, articleIDs []int64) {
	for i := range articleIDs {
		_ = r.mc.Delete("art:" + cast.FormatInt(articleIDs[i]))
	}
}
