namespace go api
include "common.thrift"
include "account.thrift"
include "archive.thrift"

struct UserProfileRequest {
    1: common.GetUserItemRequest item
}

struct UserProfileResponse {
    1: account.User user

    // published recivied comment
    // relation
    // liked
}

service APIService {
    UserProfileResponse UserProfile(1: UserProfileRequest req) (api.get="/user/profile")

}
