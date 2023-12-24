namespace go archive
include "common.thrift"

struct Article {
    1: i64 article_id
    2: i64 uid
    3: string content
}

struct CreateArticleRequest {
    1: i64 uid
    2: string content
}

struct GetArticleResponse {
    1: Article article
}

struct MGetArticlesRequest {
    1: list<i64> article_ids
}

struct MGetArticlesResponse {
    1: map<i64, Article> articles
}

struct GetUserArticlesResponse {
    1: list<Article> articles
    2: bool has_more
}

struct DeleteArticlesRequest {
    1: i64 uid
    2: list<i64> article_ids
}

service ArchiveBFF {
    common.EmptyResponse CreateArticle(1: CreateArticleRequest req)
    GetArticleResponse GetArticle(1: common.GetItemRequest req)
    MGetArticlesResponse MGetArticles(1: MGetArticlesRequest req)
    GetUserArticlesResponse GetUserArticles(1: common.GetUserItemRequest req)
    common.EmptyResponse DeleteArticles(1: DeleteArticlesRequest req)
}
