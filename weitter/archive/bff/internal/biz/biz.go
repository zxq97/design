package biz

import (
	"context"

	"github.com/google/wire"
	"github.com/zxq97/gokit/pkg/generate"
)

var ProviderSet = wire.NewSet(NewArchiveBFFUseCase)

type ArchiveBFFRepo interface {
	CreateArticle(ctx context.Context, article *Article) error
	MGetArticles(ctx context.Context, articleIDs []int64) (map[int64]*Article, error)
	GetUserArticles(ctx context.Context, uid, lastID int64, limit int8) ([]*Article, bool, error)
	DeleteArticles(ctx context.Context, articleIDs []int64) error

	GetUser(ctx context.Context, uid int64) (*User, error)
	CheckAdminUser(ctx context.Context, uid int64) (bool, error)
}

type ArchiveBFFUseCase struct {
	repo  ArchiveBFFRepo
	idGen *generate.SnowIDGen
}

func NewArchiveBFFUseCase(repo ArchiveBFFRepo) *ArchiveBFFUseCase {
	idGen, err := generate.NewSnowIDGen("archive_bff")
	if err != nil {
		panic(err)
	}
	return &ArchiveBFFUseCase{repo: repo, idGen: idGen}
}
