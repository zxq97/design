namespace go account

enum Gender {
    Undifine = 0
    Lady = 1
    Gentleman = 2
    Admin = 3
}

struct EmptyResponse {
}

struct User {
    1: i64 uid
    2: string nickname
    3: Gender gender
    4: string introduction
}

struct CreateUserRequest {
    1: User user
}

struct MGetUsersRequest {
    1: list<i64> uids
}

struct MGetUsersResponse {
    1: map<i64, User> users
}

struct DeleteUsersRequest {
    1: list<i64> uids
}

service AccountService {
    EmptyResponse CreateUser(1: CreateUserRequest req)
    MGetUsersResponse MGetUsers(1: MGetUsersRequest req)
    EmptyResponse DeleteUsers(1: DeleteUsersRequest req)
}
