package service

import (
	"context"

	"github.com/zxq97/design/fuxi/api/service/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *FuxiService) GetGenerateID(ctx context.Context, req *emptypb.Empty) (*v1.GetGenerateResponse, error) {
	gid, err := s.fu.GetGenerateID(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetGenerateResponse{Gid: gid}, nil
}
