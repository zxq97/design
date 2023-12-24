package biz

import (
	"context"

	"github.com/pkg/errors"
)

type Article struct {
	ArticleID int64
	UID       int64
	Content   string
}

type User struct {
	UID    int64
	Gender int32
}

func (uc *ArchiveBFFUseCase) CreateArticle(ctx context.Context, article *Article) error {
	// TODO check account banned?
	if _, err := uc.repo.GetUser(ctx, article.UID); err != nil {
		return err
	}
	article.ArticleID = uc.idGen.Gen()
	return uc.repo.CreateArticle(ctx, article)
}

func (uc *ArchiveBFFUseCase) GetArticle(ctx context.Context, articleID int64) (*Article, error) {
	m, err := uc.repo.MGetArticles(ctx, []int64{articleID})
	if err != nil {
		return nil, err
	}
	a, ok := m[articleID]
	if !ok {
		return nil, errors.New("archive_bff: article not found")
	}
	return a, nil
}

func (uc *ArchiveBFFUseCase) MGetArticles(ctx context.Context, articleIDs []int64) (map[int64]*Article, error) {
	return uc.repo.MGetArticles(ctx, articleIDs)
}

func (uc *ArchiveBFFUseCase) GetUserArticles(ctx context.Context, uid, lastID int64, limit int8) ([]*Article, bool, error) {
	return uc.repo.GetUserArticles(ctx, uid, lastID, limit)
}

func (uc *ArchiveBFFUseCase) DeleteArticles(ctx context.Context, uid int64, articleIDs []int64) error {
	ok, err := uc.repo.CheckAdminUser(ctx, uid)
	if err != nil {
		return err
	}
	m, err := uc.repo.MGetArticles(ctx, articleIDs)
	if err != nil {
		return err
	}
	for _, v := range m {
		if !ok && v.UID != uid {
			return errors.New("archive_bff: no permission delete")
		}
	}
	return uc.repo.DeleteArticles(ctx, articleIDs)
}
