namespace go archive

struct EmptyResponse {
}

struct Article {
    1: i64 article_id
    2: i64 uid
    3: string content
}

struct CreateArticleRequest {
    1: Article article
}

struct MGetArticlesRequest {
    1: list<i64> article_ids
}

struct MGetArticleResponse {
    1: map<i64, Article> articles
}

struct GetUserArticlesRequest {
    1: i64 uid
    2: i64 last_id
    3: i8 limit
}

struct GetUserArticlesResponse {
    1: list<Article> articles
    2: bool has_more
}

struct DelectArticlesRequest {
    1: list<i64> article_ids
}

service ArchiveService {
    EmptyResponse CreateArticle(1: CreateArticleRequest req)
    MGetArticleResponse MGetArticles(1: MGetArticlesRequest req)
    GetUserArticlesResponse GetUserArticles(1: GetUserArticlesRequest req)
    EmptyResponse DeleteArticles(1: DelectArticlesRequest req)
}
