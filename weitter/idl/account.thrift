namespace go account
include "common.thrift"

enum Gender {
    Undifine = 0
    Lady = 1
    Gentleman = 2
    Admin = 3
}

struct User {
    1: i64 uid
    2: string nickname
    3: Gender gender
    4: string introduction
}

struct CreateUserRequest {
    1: string nickname
    2: Gender gender
    3: string introduction
}

struct GetUserResponse {
    1: User user
}

struct MGetUsersResponse {
    1: map<i64, User> users
}

struct DeleteUsersRequest {
    1: i64 uid
    2: list<i64> uids
}

service AccountBFF {
    common.EmptyResponse CreateUser(1: CreateUserRequest req)
    GetUserResponse GetUser(1: common.GetItemRequest req)
    MGetUsersResponse MGetUsers(1: common.MGetItemARequest req)
    common.EmptyResponse DeleteUsers(1: DeleteUsersRequest req)
    common.CheckResponse CheckAdminUser(1: common.GetItemRequest req)
}
