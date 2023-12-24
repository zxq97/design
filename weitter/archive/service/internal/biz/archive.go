package biz

import (
	"context"

	"github.com/zxq97/design/weitter/archive/pkg/model"
)

func (uc *ArchiveUseCase) CreateArticle(ctx context.Context, article *model.Article) error {
	return uc.repo.CreateArticle(ctx, article)
}

func (uc *ArchiveUseCase) MGetArticles(ctx context.Context, articleIDs []int64) (map[int64]*model.Article, error) {
	return uc.repo.MGetArticles(ctx, articleIDs)
}

func (uc *ArchiveUseCase) GetUserArticles(ctx context.Context, uid, lastID int64, limit int8) ([]*model.Article, bool, error) {
	return uc.repo.GetUserArticles(ctx, uid, lastID, limit)
}

func (uc *ArchiveUseCase) DeleteArticles(ctx context.Context, articleIDs []int64) error {
	return uc.repo.DeleteArticles(ctx, articleIDs)
}
