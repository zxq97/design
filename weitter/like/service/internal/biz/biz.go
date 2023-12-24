package biz

import (
	"context"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewLikeUseCase)

type LikeRepo interface {
	Like(ctx context.Context, objID, uid, author int64, objType int32) error
	Unlike(ctx context.Context, objID, uid int64, objType int32) error
	GetLikedUsers(ctx context.Context, objID, lastUID int64, objType int32, limit int8, preload bool) ([]int64, bool, error)
	//GetRcvLikedRecord(ctx context.Context, author, objID, uid int64, objType int32, limit int8) ([]*model.Like, bool, error)
	//GetPubLikedRecord(ctx context.Context, uid, objID, author int64, objType int32, limit int8) ([]*model.Like, bool, error)
	MGetLikedState(ctx context.Context, m map[int32][]int64, uid int64) (map[int32]map[int64]bool, error)
	MGetLikedCount(ctx context.Context, m map[int32][]int64) (map[int32]map[int64]int32, error)

	MUpdateLikesCount(ctx context.Context, m map[int32]map[int64]int32) error
}

type LikeUseCase struct {
	repo LikeRepo
}

func NewLikeUseCase(repo LikeRepo) *LikeUseCase {
	return &LikeUseCase{repo: repo}
}
