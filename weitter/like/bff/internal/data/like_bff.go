package data

import (
	"context"

	"github.com/zxq97/design/weitter/kitex_gen/account/accountbff"
	"github.com/zxq97/design/weitter/kitex_gen/common"
	"github.com/zxq97/design/weitter/like/bff/internal/biz"
	"github.com/zxq97/design/weitter/like/kitex_gen/like"
	"github.com/zxq97/design/weitter/like/kitex_gen/like/likeservice"
)

var _ biz.LikeBFFRepo = (*likeBFFRepo)(nil)

type likeBFFRepo struct {
	likeClient    likeservice.Client
	accountClient accountbff.Client
}

func NewLikeBFFRepo(likeClient likeservice.Client, accountClient accountbff.Client) biz.LikeBFFRepo {
	return &likeBFFRepo{likeClient: likeClient, accountClient: accountClient}
}

func (r *likeBFFRepo) Like(ctx context.Context, p *biz.Like) error {
	_, err := r.likeClient.Like(ctx, &like.LikeRequest{
		ObjId:   p.ObjID,
		ObjType: p.ObjType,
		Uid:     p.UID,
		Author:  p.Author,
	})
	return err
}

func (r *likeBFFRepo) Unlike(ctx context.Context, p *biz.Like) error {
	_, err := r.likeClient.Unlike(ctx, &like.LikeRequest{
		ObjId:   p.ObjID,
		ObjType: p.ObjType,
		Uid:     p.UID,
		Author:  p.Author,
	})
	return err
}

func (r *likeBFFRepo) GetLikedUsers(ctx context.Context, objID, lastUID int64, objType int32, limit int8, preload bool) ([]int64, bool, error) {
	res, err := r.likeClient.GetLikedUsers(ctx, &like.GetLikedUsersRequest{
		ObjId:   objID,
		ObjType: objType,
		LastUid: lastUID,
		Limit:   limit,
		Preload: preload,
	})
	return res.Uids, res.HasMore, err
}

func (r *likeBFFRepo) GetRcvLikedList(ctx context.Context, p *biz.Like, limit int8) ([]*biz.Like, bool, error) {
	res, err := r.likeClient.GetRcvLikedList(ctx, &like.GetLikedRecordRequest{
		ObjId:   p.ObjID,
		ObjType: p.ObjType,
		Uid:     p.UID,
		Author:  p.Author,
		Limit:   limit,
	})
	if err != nil {
		return nil, false, err
	}
	ls := make([]*biz.Like, len(res.Likes))
	for i := range res.Likes {
		ls[i] = &biz.Like{
			ObjID:   res.Likes[i].ObjId,
			ObjType: res.Likes[i].ObjType,
			UID:     res.Likes[i].Uid,
		}
	}
	return ls, res.HasMore, nil
}

func (r *likeBFFRepo) GetPubLikedList(ctx context.Context, p *biz.Like, limit int8) ([]*biz.Like, bool, error) {
	res, err := r.likeClient.GetPubLikedList(ctx, &like.GetLikedRecordRequest{
		ObjId:   p.ObjID,
		ObjType: p.ObjType,
		Uid:     p.UID,
		Author:  p.Author,
		Limit:   limit,
	})
	if err != nil {
		return nil, false, err
	}
	ls := make([]*biz.Like, len(res.Likes))
	for i := range res.Likes {
		ls[i] = &biz.Like{
			ObjID:   res.Likes[i].ObjId,
			ObjType: res.Likes[i].ObjType,
			UID:     res.Likes[i].Uid,
		}
	}
	return ls, res.HasMore, nil
}

func (r *likeBFFRepo) MGetLikedState(ctx context.Context, m map[int32][]int64, uid int64) (map[int32]map[int64]bool, error) {
	res, err := r.likeClient.MGetLikedState(ctx, &like.MGetLikedStateRequest{
		Obj: m,
		Uid: uid,
	})
	return res.M, err
}

func (r *likeBFFRepo) MGetLikedCount(ctx context.Context, m map[int32][]int64) (map[int32]map[int64]int32, error) {
	res, err := r.likeClient.MGetLikedCount(ctx, &like.MGetLikedCountRequest{
		Obj: m,
	})
	return res.M, err
}

func (r *likeBFFRepo) MUpdateLikesCount(ctx context.Context, m map[int32]map[int64]int32) error {
	_, err := r.likeClient.MUpdateLikesCount(ctx, &like.UpdateLikesCountRequest{
		M: m,
	})
	return err
}

func (r *likeBFFRepo) CheckAdminUser(ctx context.Context, uid int64) (bool, error) {
	res, err := r.accountClient.CheckAdminUser(ctx, &common.GetItemRequest{Id: uid})
	if err != nil {
		return false, err
	}
	return res.Ok, nil
}
