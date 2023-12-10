package biz

import "context"

type FuxiBFFRepo interface {
	Set(ctx context.Context, realUrl, shortUrl string) error
	Get(ctx context.Context, shortUrl string) (string, error)
	Allocation(ctx context.Context, realUrl string) (string, error)
}

type FuxiBFFUseCase struct {
	repo FuxiBFFRepo
}

func NewFuxiBFFUseCase(repo FuxiBFFRepo) *FuxiBFFUseCase {
	return &FuxiBFFUseCase{repo: repo}
}

func (uc *FuxiBFFUseCase) SetUrl(ctx context.Context, realUrl, shortUrl string) error {
	return uc.repo.Set(ctx, realUrl, shortUrl)
}

func (uc *FuxiBFFUseCase) GetUrl(ctx context.Context, shortUrl string) (string, error) {
	return uc.repo.Get(ctx, shortUrl)
}

func (uc *FuxiBFFUseCase) AllocationUrl(ctx context.Context, realUrl string) (string, error) {
	return uc.repo.Allocation(ctx, realUrl)
}
