package data

import (
	"context"

	"github.com/zxq97/design/fuxi/pkg/constant"
	"github.com/zxq97/design/fuxi/pkg/query"
	"github.com/zxq97/design/fuxi/service/internal/biz"
	"gorm.io/gorm"
)

var _ biz.FuxiRepo = (*fuxiRepo)(nil)

type fuxiRepo struct {
	q *query.Query
}

func NewFuxiRepo(db *gorm.DB) biz.FuxiRepo {
	return &fuxiRepo{q: query.Use(db)}
}

func (r *fuxiRepo) Load(ctx context.Context) ([]int64, error) {
	rows, err := r.q.WithContext(ctx).URLMap.Select(r.q.URLMap.ID).Where(r.q.URLMap.Status.Eq(constant.URLMapStatusUnused)).Order(r.q.URLMap.ID.Asc()).Limit(constant.ReloadSize).Find()
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0, len(rows))

	for i := range rows {
		ids = append(ids, rows[i].ID)
	}
	return ids, nil
}
