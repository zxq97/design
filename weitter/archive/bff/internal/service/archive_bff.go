package service

import (
	"context"

	"github.com/zxq97/design/weitter/archive/bff/internal/biz"
	"github.com/zxq97/design/weitter/kitex_gen/archive"
	"github.com/zxq97/design/weitter/kitex_gen/common"
)

// CreateArticle implements the ArchiveBFFImpl interface.
func (s *ArchiveBFFImpl) CreateArticle(ctx context.Context, req *archive.CreateArticleRequest) (resp *common.EmptyResponse, err error) {
	resp = new(common.EmptyResponse)
	err = s.au.CreateArticle(ctx, &biz.Article{
		UID:     req.Uid,
		Content: req.Content,
	})
	return
}

// GetArticle implements the ArchiveBFFImpl interface.
func (s *ArchiveBFFImpl) GetArticle(ctx context.Context, req *common.GetItemRequest) (resp *archive.GetArticleResponse, err error) {
	resp = new(archive.GetArticleResponse)
	a, err := s.au.GetArticle(ctx, req.Id)
	if err != nil {
		return resp, err
	}
	resp.Article = &archive.Article{
		ArticleId: a.ArticleID,
		Uid:       a.UID,
		Content:   a.Content,
	}
	return
}

// MGetArticles implements the ArchiveBFFImpl interface.
func (s *ArchiveBFFImpl) MGetArticles(ctx context.Context, req *archive.MGetArticlesRequest) (resp *archive.MGetArticlesResponse, err error) {
	resp = new(archive.MGetArticlesResponse)
	m, err := s.au.MGetArticles(ctx, req.ArticleIds)
	if err != nil {
		return resp, err
	}
	resp.Articles = make(map[int64]*archive.Article, len(m))
	for k, v := range m {
		resp.Articles[k] = &archive.Article{
			ArticleId: v.ArticleID,
			Uid:       v.UID,
			Content:   v.Content,
		}
	}
	return
}

// GetUserArticles implements the ArchiveBFFImpl interface.
func (s *ArchiveBFFImpl) GetUserArticles(ctx context.Context, req *common.GetUserItemRequest) (resp *archive.GetUserArticlesResponse, err error) {
	resp = new(archive.GetUserArticlesResponse)
	as, hasMore, err := s.au.GetUserArticles(ctx, req.Uid, req.LastId, req.Limit)
	if err != nil {
		return resp, err
	}
	resp.Articles = make([]*archive.Article, len(as))
	for i := range as {
		resp.Articles[i] = &archive.Article{
			ArticleId: as[i].ArticleID,
			Uid:       as[i].UID,
			Content:   as[i].Content,
		}
	}
	resp.HasMore = hasMore
	return
}

// DeleteArticles implements the ArchiveBFFImpl interface.
func (s *ArchiveBFFImpl) DeleteArticles(ctx context.Context, req *archive.DeleteArticlesRequest) (resp *common.EmptyResponse, err error) {
	resp = new(common.EmptyResponse)
	err = s.au.DeleteArticles(ctx, req.Uid, req.ArticleIds)
	return
}
