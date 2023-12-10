package data

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	v1 "github.com/zxq97/design/fuxi/api/service/v1"
	"github.com/zxq97/design/fuxi/bff/internal/biz"
	"github.com/zxq97/design/fuxi/pkg/constant"
	"github.com/zxq97/design/fuxi/pkg/model"
	"github.com/zxq97/design/fuxi/pkg/query"
	"github.com/zxq97/gokit/pkg/cast"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

var _ biz.FuxiBFFRepo = (*fuxiBFFRepo)(nil)

type fuxiBFFRepo struct {
	redis *redis.Client
	q     *query.Query
	fc    v1.FuxiClient
}

func NewFuxiBFFRepo(redis *redis.Client, db *gorm.DB, fc v1.FuxiClient) biz.FuxiBFFRepo {
	return &fuxiBFFRepo{redis: redis, q: query.Use(db), fc: fc}
}

func (r *fuxiBFFRepo) Set(ctx context.Context, realUrl, shortUrl string) error {
	bs, err := base64.StdEncoding.DecodeString(shortUrl)
	if err != nil {
		return err
	}
	gid := cast.ParseInt(string(bs), 0)

	row, err := r.q.WithContext(ctx).URLMap.Select(r.q.URLMap.ID, r.q.URLMap.ShortURL, r.q.URLMap.RealURL, r.q.URLMap.Status).Where(r.q.URLMap.ShortURL.Eq(shortUrl)).Or(r.q.URLMap.ID.Eq(gid)).Or(r.q.URLMap.RealURL.Eq(realUrl)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if row != nil {
		if row.RealURL == realUrl {
			if row.ShortURL != shortUrl {
				return errors.New("fuxi: real url was used")
			}
		} else if row.ShortURL == shortUrl {
			if row.RealURL != realUrl {
				return errors.New("fuxi: short url was used")
			}
			return nil
		} else if row.ID == gid {
			if row.ShortURL != shortUrl || row.RealURL != realUrl {
				return errors.New("fuxi: url generate id was used")
			}
		}
		if row.Status == constant.URLMapStatusUsed {
			return nil
		}
		row.Status = constant.URLMapStatusUsed
		return r.q.WithContext(ctx).URLMap.Save(row)
	}

	if err = r.q.WithContext(ctx).URLMap.Create(&model.URLMap{
		ID:       gid,
		ShortURL: shortUrl,
		RealURL:  realUrl,
		Status:   constant.URLMapStatusUsed,
	}); err != nil {
		return err
	}

	key := fmt.Sprintf("fuxi_surl:%s", shortUrl)
	_ = r.redis.Set(ctx, key, realUrl, time.Hour*24*7).Err()
	return nil
}

func (r *fuxiBFFRepo) Get(ctx context.Context, shortUrl string) (string, error) {
	key := fmt.Sprintf("fuxi_surl:%s", shortUrl)
	url, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		return url, nil
	}

	row, err := r.q.WithContext(ctx).URLMap.Select(r.q.URLMap.RealURL).Where(r.q.URLMap.ShortURL.Eq(shortUrl), r.q.URLMap.Status.Eq(constant.URLMapStatusUsed)).First()
	if err != nil {
		return "", err
	}

	_ = r.redis.Set(ctx, key, row.RealURL, time.Hour*24).Err()
	return row.RealURL, nil
}

func (r *fuxiBFFRepo) Allocation(ctx context.Context, realUrl string) (string, error) {
	res, err := r.fc.GetGenerateID(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}

	shortUrl := base64.StdEncoding.EncodeToString([]byte(cast.FormatInt(res.Gid)))

	row, err := r.q.WithContext(ctx).URLMap.Where(r.q.URLMap.ID.Eq(res.Gid), r.q.URLMap.ShortURL.Eq(shortUrl), r.q.URLMap.Status.Eq(constant.URLMapStatusUnused)).First()
	if err != nil {
		return "", err
	}

	row.Status = constant.URLMapStatusUsed
	row.RealURL = realUrl
	if err = r.q.WithContext(ctx).URLMap.Save(row); err != nil {
		return "", err
	}

	key := fmt.Sprintf("fuxi_surl:%s", shortUrl)
	_ = r.redis.Set(ctx, key, realUrl, time.Hour*24*7).Err()
	return shortUrl, nil
}
