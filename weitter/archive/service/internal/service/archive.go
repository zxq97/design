package service

import (
	"context"

	"github.com/zxq97/design/weitter/archive/kitex_gen/archive"
	"github.com/zxq97/design/weitter/archive/pkg/model"
)

// CreateArticle implements the ArchiveServiceImpl interface.
func (s *ArchiveServiceImpl) CreateArticle(ctx context.Context, req *archive.CreateArticleRequest) (resp *archive.EmptyResponse, err error) {
	resp = new(archive.EmptyResponse)
	err = s.au.CreateArticle(ctx, &model.Article{
		ArticleID: req.Article.ArticleId,
		UID:       req.Article.Uid,
		Content:   req.Article.Content,
	})
	return
}

// MGetArticles implements the ArchiveServiceImpl interface.
func (s *ArchiveServiceImpl) MGetArticles(ctx context.Context, req *archive.MGetArticlesRequest) (resp *archive.MGetArticleResponse, err error) {
	resp = new(archive.MGetArticleResponse)
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

// GetUserArticles implements the ArchiveServiceImpl interface.
func (s *ArchiveServiceImpl) GetUserArticles(ctx context.Context, req *archive.GetUserArticlesRequest) (resp *archive.GetUserArticlesResponse, err error) {
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

// DeleteArticles implements the ArchiveServiceImpl interface.
func (s *ArchiveServiceImpl) DeleteArticles(ctx context.Context, req *archive.DelectArticlesRequest) (resp *archive.EmptyResponse, err error) {
	resp = new(archive.EmptyResponse)
	err = s.au.DeleteArticles(ctx, req.ArticleIds)
	return
}
