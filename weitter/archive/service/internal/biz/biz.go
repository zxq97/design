package biz

import (
	"context"

	"github.com/google/wire"
	"github.com/zxq97/design/weitter/archive/pkg/model"
)

var ProviderSet = wire.NewSet(NewArchiveUseCase)

type ArchiveRepo interface {
	CreateArticle(ctx context.Context, article *model.Article) error
	MGetArticles(ctx context.Context, articleIDs []int64) (map[int64]*model.Article, error)
	GetUserArticles(ctx context.Context, uid, lastID int64, limit int8) ([]*model.Article, bool, error)
	DeleteArticles(ctx context.Context, articleIDs []int64) error
}

type ArchiveUseCase struct {
	repo ArchiveRepo
}

func NewArchiveUseCase(repo ArchiveRepo) *ArchiveUseCase {
	return &ArchiveUseCase{repo: repo}
}
