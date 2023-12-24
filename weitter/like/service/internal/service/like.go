package service

import (
	"context"

	like "github.com/zxq97/design/weitter/like/kitex_gen/like"
)

// Like implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) Like(ctx context.Context, req *like.LikeRequest) (resp *like.EmptyResponse, err error) {
	resp = new(like.EmptyResponse)
	err = s.lu.Like(ctx, req.ObjId, req.Uid, req.Author, req.ObjType)
	return
}

// Unlike implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) Unlike(ctx context.Context, req *like.LikeRequest) (resp *like.EmptyResponse, err error) {
	resp = new(like.EmptyResponse)
	err = s.lu.Unlike(ctx, req.ObjId, req.Uid, req.ObjType)
	return
}

// GetLikedUsers implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) GetLikedUsers(ctx context.Context, req *like.GetLikedUsersRequest) (resp *like.GetLikedUsersReponse, err error) {
	resp = new(like.GetLikedUsersReponse)
	resp.Uids, resp.HasMore, err = s.lu.GetLikedUsers(ctx, req.ObjId, req.LastUid, req.ObjType, req.Limit, req.Preload)
	return
}

// GetRcvLikedList implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) GetRcvLikedList(ctx context.Context, req *like.GetLikedRecordRequest) (resp *like.GetLikedRecordResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPubLikedList implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) GetPubLikedList(ctx context.Context, req *like.GetLikedRecordRequest) (resp *like.GetLikedRecordResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetLikedState implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) MGetLikedState(ctx context.Context, req *like.MGetLikedStateRequest) (resp *like.MGetLikedStateResponse, err error) {
	resp = new(like.MGetLikedStateResponse)
	resp.M, err = s.lu.MGetLikedState(ctx, req.Obj, req.Uid)
	return
}

// MGetLikedCount implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) MGetLikedCount(ctx context.Context, req *like.MGetLikedCountRequest) (resp *like.MGetLikedCountResponse, err error) {
	resp = new(like.MGetLikedCountResponse)
	resp.M, err = s.lu.MGetLikedCount(ctx, req.Obj)
	return
}

// MUpdateLikesCount implements the LikeServiceImpl interface.
func (s *LikeServiceImpl) MUpdateLikesCount(ctx context.Context, req *like.UpdateLikesCountRequest) (resp *like.EmptyResponse, err error) {
	resp = new(like.EmptyResponse)
	err = s.lu.MUpdateLikesCount(ctx, req.M)
	return
}
