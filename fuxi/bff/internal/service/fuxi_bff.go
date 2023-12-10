package service

import (
	"context"

	"github.com/zxq97/design/fuxi/api/bff/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *FuxiBFFService) SetUrl(ctx context.Context, req *v1.SetUrlRequest) (*emptypb.Empty, error) {
	if err := s.fu.SetUrl(ctx, req.RealUrl, req.ShortUrl); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *FuxiBFFService) GetUrl(ctx context.Context, req *v1.GetUrlRequest) (*v1.GetUrlResponse, error) {
	url, err := s.fu.GetUrl(ctx, req.ShortUrl)
	if err != nil {
		return nil, err
	}
	return &v1.GetUrlResponse{RealUrl: url}, nil
}

func (s *FuxiBFFService) AllocationUrl(ctx context.Context, req *v1.AllocationRequest) (*v1.AllocationResponse, error) {
	url, err := s.fu.AllocationUrl(ctx, req.RealUrl)
	if err != nil {
		return nil, err
	}
	return &v1.AllocationResponse{ShortUrl: url}, nil
}
