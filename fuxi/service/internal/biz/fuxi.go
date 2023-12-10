package biz

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zxq97/design/fuxi/pkg/constant"
)

type FuxiRepo interface {
	Load(ctx context.Context) ([]int64, error)
}

type FuxiUseCase struct {
	repo FuxiRepo
	lock sync.Mutex
	ch   chan int64
	cnt  int32
}

func NewFuxiUseCase(repo FuxiRepo) *FuxiUseCase {
	return &FuxiUseCase{repo: repo, ch: make(chan int64, constant.MemorySize)}
}

func (uc *FuxiUseCase) Load(unlock bool) error {
	if unlock {
		defer uc.lock.Unlock()
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	ids, err := uc.repo.Load(ctx)
	if err != nil {
		return err
	}

	for i := range ids {
		uc.ch <- ids[i]
		atomic.AddInt32(&uc.cnt, 1)
	}
	return nil
}

func (uc *FuxiUseCase) GetGenerateID(ctx context.Context) (int64, error) {
	select {
	case id := <-uc.ch:
		l := atomic.AddInt32(&uc.cnt, -1)
		if l <= constant.LoadWarnSize && uc.lock.TryLock() {
			go uc.Load(true)
		}
		return id, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}
