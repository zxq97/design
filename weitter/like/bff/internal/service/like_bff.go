package service

import (
	"context"

	"github.com/zxq97/design/weitter/kitex_gen/common"
	"github.com/zxq97/design/weitter/kitex_gen/like"
)

// Like implements the LikeBFFImpl interface.
func (s *LikeBFFImpl) Like(ctx context.Context, req *like.LikeRequest) (resp *common.EmptyResponse, err error) {
	// TODO: Your code here...
	return
}

// Unlike implements the LikeBFFImpl interface.
func (s *LikeBFFImpl) Unlike(ctx context.Context, req *like.LikeRequest) (resp *common.EmptyResponse, err error) {
	// TODO: Your code here...
	return
}

// GetLikedUsers implements the LikeBFFImpl interface.
func (s *LikeBFFImpl) GetLikedUsers(ctx context.Context, req *like.GetLikedUsersRequest) (resp *like.GetLikedUsersResponse, err error) {
	// TODO: Your code here...
	return
}

// GetRcvLikedList implements the LikeBFFImpl interface.
func (s *LikeBFFImpl) GetRcvLikedList(ctx context.Context, req *like.GetLikedRecordRequest) (resp *like.GetLikedRecordResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPubLikedList implements the LikeBFFImpl interface.
func (s *LikeBFFImpl) GetPubLikedList(ctx context.Context, req *like.GetLikedRecordRequest) (resp *like.GetLikedRecordResponse, err error) {
	// TODO: Your code here...
	return
}

// GetLikedState implements the LikeBFFImpl interface.
func (s *LikeBFFImpl) GetLikedState(ctx context.Context, req *like.GetLikedStateRequest) (resp *like.GetLikedStateResponse, err error) {
	// TODO: Your code here...
	return
}

// GetLikedCount implements the LikeBFFImpl interface.
func (s *LikeBFFImpl) GetLikedCount(ctx context.Context, req *like.GetLikedCountRequest) (resp *like.GetLikedCountResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetLikedState implements the LikeBFFImpl interface.
func (s *LikeBFFImpl) MGetLikedState(ctx context.Context, req *like.MGetLikedStateRequest) (resp *like.MGetLikedStateResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetLikedCount implements the LikeBFFImpl interface.
func (s *LikeBFFImpl) MGetLikedCount(ctx context.Context, req *like.MGetLikedCountRequest) (resp *like.MGetLikedCountResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateLikeCount implements the LikeBFFImpl interface.
func (s *LikeBFFImpl) UpdateLikeCount(ctx context.Context, req *like.UpdateLikeCountRequest) (resp *common.EmptyResponse, err error) {
	// TODO: Your code here...
	return
}

// MUpdateLikesCount implements the LikeBFFImpl interface.
func (s *LikeBFFImpl) MUpdateLikesCount(ctx context.Context, req *like.MUpdateLikesCountRequest) (resp *common.EmptyResponse, err error) {
	// TODO: Your code here...
	return
}
