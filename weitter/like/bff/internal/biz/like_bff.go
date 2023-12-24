package biz

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zxq97/design/weitter/like/pkg/constant"
)

type Like struct {
	ObjID   int64
	ObjType int32
	UID     int64
	Author  int64
}

func (uc *LikeBFFUseCase) Like(ctx context.Context, p *Like) error {
	ok, err := uc.GetLikedState(ctx, p.ObjID, p.UID, p.ObjType)
	if err != nil || ok {
		return err
	}
	return uc.repo.Like(ctx, p)
}

func (uc *LikeBFFUseCase) Unlike(ctx context.Context, p *Like) error {
	ok, err := uc.GetLikedState(ctx, p.ObjID, p.UID, p.ObjType)
	if err != nil || !ok {
		return err
	}
	return uc.repo.Unlike(ctx, p)
}

func (uc *LikeBFFUseCase) GetLikedUsers(ctx context.Context, objID int64, objType int32) ([]int64, error) {
	// TODO switch etcd get preload and limit
	uids, _, err := uc.repo.GetLikedUsers(ctx, objID, 0, objType, constant.LikedUsersCount, false)
	return uids, err
}

func (uc *LikeBFFUseCase) GetRcvLikedList(ctx context.Context, p *Like, operatorUID int64, limit int8) ([]*Like, bool, error) {
	if p.Author != operatorUID {
		return nil, false, errors.New("like_bff: no permission browse")
	}
	return uc.repo.GetRcvLikedList(ctx, p, limit)
}

func (uc *LikeBFFUseCase) GetPubLikedList(ctx context.Context, p *Like, operatorUID int64, limit int8) ([]*Like, bool, error) {
	if p.UID != operatorUID {
		return nil, false, errors.New("like_bff: no permission browse")
	}
	return uc.repo.GetPubLikedList(ctx, p, limit)
}

func (uc *LikeBFFUseCase) GetLikedState(ctx context.Context, objID, uid int64, objType int32) (bool, error) {
	m, err := uc.repo.MGetLikedState(ctx, map[int32][]int64{objType: []int64{objID}}, uid)
	if err != nil {
		return false, err
	}
	return m[objType][objID], nil
}

func (uc *LikeBFFUseCase) GetLikedCount(ctx context.Context, objID int64, objType int32) (int32, error) {
	m, err := uc.repo.MGetLikedCount(ctx, map[int32][]int64{objType: []int64{objID}})
	if err != nil {
		return 0, err
	}
	return m[objType][objID], nil
}

func (uc *LikeBFFUseCase) MGetLikedState(ctx context.Context, m map[int32][]int64, uid int64) (map[int32]map[int64]bool, error) {
	return uc.repo.MGetLikedState(ctx, m, uid)
}

func (uc *LikeBFFUseCase) MGetLikedCount(ctx context.Context, m map[int32][]int64) (map[int32]map[int64]int32, error) {
	return uc.repo.MGetLikedCount(ctx, m)
}

func (uc *LikeBFFUseCase) UpdateLikedCount(ctx context.Context, objID, uid int64, objType, count int32) error {
	return uc.MUpdateLikedCount(ctx, map[int32]map[int64]int32{objType: {objID: count}}, uid)
}

func (uc *LikeBFFUseCase) MUpdateLikedCount(ctx context.Context, m map[int32]map[int64]int32, uid int64) error {
	ok, err := uc.repo.CheckAdminUser(ctx, uid)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("like_bff: no permission update count")
	}
	return uc.repo.MUpdateLikesCount(ctx, m)
}
