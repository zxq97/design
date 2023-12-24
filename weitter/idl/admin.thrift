namespace go admin
include "common.thrift"
include "account.thrift"
include "archive.thrift"
include "like.thrift"

service AdminService {
//    like
    common.EmptyResponse UpdateLikedCount(1: like.UpdateLikeCountRequest req) (api.post="/like/update/count")
    common.EmptyResponse MUpdateLikedCount(1: like.MUpdateLikesCountRequest req) (api.post="/like/multi_update/count")
//     TODO user archive comment list like

//    archive
    common.EmptyResponse DeleteArticles(1: archive.DeleteArticlesRequest req) (api.post="/archive/delete")
//     TODO user archive list, time desc list

//    account
    common.EmptyResponse DeleteUsers(1: account.DeleteUsersRequest req) (api.post="/account/delete")
    common.EmptyResponse CreateAdminUser(1: account.CreateUserRequest req) (api.post="/account/create")

//    comment

//    relation
}
