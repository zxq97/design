package biz

import (
	"context"

	"github.com/zxq97/design/weitter/like/pkg/model"
)

func (uc *LikeUseCase) Like(ctx context.Context, objID, uid, author int64, objType int32) error {
	return uc.repo.Like(ctx, objID, uid, author, objType)
}

func (uc *LikeUseCase) Unlike(ctx context.Context, objID, uid int64, objType int32) error {
	return uc.repo.Unlike(ctx, objID, uid, objType)
}

func (uc *LikeUseCase) GetLikedUsers(ctx context.Context, objID, lastUID int64, objType int32, limit int8, preload bool) ([]int64, bool, error) {
	uids, hasMore, err := uc.repo.GetLikedUsers(ctx, objID, lastUID, objType, limit, preload)
	if err != nil {
		return nil, false, err
	}
	if !preload {
		hasMore = false
	}
	return uids, hasMore, nil
}

func (uc *LikeUseCase) GetRcvLikedRecord(ctx context.Context, objID, author, uid int64, objType int32, limit int8) (*[]model.Like, bool, error) {
	// TODO
	return nil, false, nil
}

func (uc *LikeUseCase) GetPubLikedRecord(ctx context.Context, objID, uid int64, objType int32, limit int8) ([]*model.Like, bool, error) {
	// TODO
	return nil, false, nil
}

func (uc *LikeUseCase) MGetLikedState(ctx context.Context, m map[int32][]int64, uid int64) (map[int32]map[int64]bool, error) {
	return uc.repo.MGetLikedState(ctx, m, uid)
}

func (uc *LikeUseCase) MGetLikedCount(ctx context.Context, m map[int32][]int64) (map[int32]map[int64]int32, error) {
	return uc.repo.MGetLikedCount(ctx, m)
}

func (uc *LikeUseCase) MUpdateLikesCount(ctx context.Context, m map[int32]map[int64]int32) error {
	return uc.repo.MUpdateLikesCount(ctx, m)
}
