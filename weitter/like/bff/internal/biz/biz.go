package biz

import (
	"context"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewLikeBFFUseCase)

type LikeBFFRepo interface {
	Like(ctx context.Context, like *Like) error
	Unlike(ctx context.Context, like *Like) error
	GetLikedUsers(ctx context.Context, objID, lastUID int64, objType int32, limit int8, preload bool) ([]int64, bool, error)
	GetRcvLikedList(ctx context.Context, like *Like, limit int8) ([]*Like, bool, error)
	GetPubLikedList(ctx context.Context, like *Like, limit int8) ([]*Like, bool, error)
	MGetLikedState(ctx context.Context, m map[int32][]int64, uid int64) (map[int32]map[int64]bool, error)
	MGetLikedCount(ctx context.Context, m map[int32][]int64) (map[int32]map[int64]int32, error)

	MUpdateLikesCount(ctx context.Context, m map[int32]map[int64]int32) error

	CheckAdminUser(ctx context.Context, uid int64) (bool, error)
}

type LikeBFFUseCase struct {
	repo LikeBFFRepo
}

func NewLikeBFFUseCase(repo LikeBFFRepo) *LikeBFFUseCase {
	return &LikeBFFUseCase{repo: repo}
}
