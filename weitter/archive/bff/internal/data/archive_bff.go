package data

import (
	"context"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/pkg/errors"
	"github.com/zxq97/design/weitter/archive/bff/internal/biz"
	"github.com/zxq97/design/weitter/archive/kitex_gen/archive"
	"github.com/zxq97/design/weitter/archive/kitex_gen/archive/archiveservice"
	"github.com/zxq97/design/weitter/kitex_gen/account/accountbff"
	"github.com/zxq97/design/weitter/kitex_gen/common"
	"github.com/zxq97/gokit/pkg/cast"
)

var _ biz.ArchiveBFFRepo = (*archiveBFFRepo)(nil)

type archiveBFFRepo struct {
	mc            *memcache.Client
	archiveClient archiveservice.Client
	accountClient accountbff.Client
}

func NewArchiveBFFRepo(mc *memcache.Client, archiveClient archiveservice.Client) biz.ArchiveBFFRepo {
	return &archiveBFFRepo{mc: mc, archiveClient: archiveClient}
}

func (r *archiveBFFRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	res, _ := r.mc.Get("pub_arc" + cast.FormatInt(article.UID))
	if res != nil && string(res.Value) == article.Content {
		return errors.New("archive_bff: publish content dup")
	}
	if _, err := r.archiveClient.CreateArticle(ctx, &archive.CreateArticleRequest{
		Article: &archive.Article{
			ArticleId: article.ArticleID,
			Uid:       article.UID,
			Content:   article.Content,
		},
	}); err != nil {
		return err
	}
	_ = r.mc.Set(&memcache.Item{Key: "pub_arc" + cast.FormatInt(article.UID), Value: []byte(article.Content), Expiration: 60 * 3})
	return nil
}

func (r *archiveBFFRepo) MGetArticles(ctx context.Context, articleIDs []int64) (map[int64]*biz.Article, error) {
	res, err := r.archiveClient.MGetArticles(ctx, &archive.MGetArticlesRequest{
		ArticleIds: articleIDs,
	})
	if err != nil {
		return nil, err
	}
	m := make(map[int64]*biz.Article, len(articleIDs))
	for k, v := range res.Articles {
		m[k] = &biz.Article{
			ArticleID: v.ArticleId,
			UID:       v.Uid,
			Content:   v.Content,
		}
	}
	return m, nil
}

func (r *archiveBFFRepo) GetUserArticles(ctx context.Context, uid, lastID int64, limit int8) ([]*biz.Article, bool, error) {
	res, err := r.archiveClient.GetUserArticles(ctx, &archive.GetUserArticlesRequest{
		Uid:    uid,
		LastId: lastID,
		Limit:  limit,
	})
	if err != nil {
		return nil, false, err
	}
	as := make([]*biz.Article, len(res.Articles))
	for i := range res.Articles {
		as[i] = &biz.Article{
			ArticleID: res.Articles[i].ArticleId,
			UID:       res.Articles[i].Uid,
			Content:   res.Articles[i].Content,
		}
	}
	return as, res.HasMore, nil
}

func (r *archiveBFFRepo) DeleteArticles(ctx context.Context, articleIDs []int64) error {
	_, err := r.archiveClient.DeleteArticles(ctx, &archive.DelectArticlesRequest{
		ArticleIds: articleIDs,
	})
	return err
}

func (r *archiveBFFRepo) GetUser(ctx context.Context, uid int64) (*biz.User, error) {
	res, err := r.accountClient.GetUser(ctx, &common.GetItemRequest{Id: uid})
	if err != nil {
		return nil, err
	}
	u := &biz.User{
		UID:    res.User.Uid,
		Gender: int32(res.User.Gender),
	}
	return u, nil
}

func (r *archiveBFFRepo) CheckAdminUser(ctx context.Context, uid int64) (bool, error) {
	res, err := r.accountClient.CheckAdminUser(ctx, &common.GetItemRequest{Id: uid})
	if err != nil {
		return false, err
	}
	return res.Ok, nil
}
