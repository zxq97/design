namespace go like

struct EmptyResponse {
}

struct LikeItem {
    1: i64 obj_id
    2: i32 obj_type
    3: i64 uid
}

struct LikeRequest {
    1: i64 obj_id
    2: i32 obj_type
    3: i64 uid
    4: i64 author
}

struct GetLikedUsersRequest {
    1: i64 obj_id
    2: i32 obj_type
    3: i64 last_uid
    4: i8 limit
    5: bool preload
}

// only limit
struct GetLikedUsersReponse {
    1: list<i64> uids
    2: bool has_more
}

struct GetLikedRecordRequest {
    1: i64 obj_id
    2: i32 obj_type
    3: i64 uid
    4: i64 author
    5: i8 limit
}

struct GetLikedRecordResponse {
    1: list<LikeItem> likes
    2: bool has_more
}

struct MGetLikedStateRequest {
    1: map<i32, list<i64>> obj // k obj_type v obj_ids
    2: i64 uid
}

struct MGetLikedStateResponse {
    1: map<i32, map<i64, bool>> m
}

struct MGetLikedCountRequest {
    1: map<i32, list<i64>> obj // k obj_type v obj_ids
}

struct MGetLikedCountResponse {
    1: map<i32, map<i64, i32>> m
}

struct UpdateLikesCountRequest {
    1: map<i32, map<i64, i32>> m // k obk_type vk obj_id vv count
}

service LikeService {
    EmptyResponse Like(1: LikeRequest req)
    EmptyResponse Unlike(1: LikeRequest req)
    GetLikedUsersReponse GetLikedUsers(1: GetLikedUsersRequest req)
    GetLikedRecordResponse GetRcvLikedList(1: GetLikedRecordRequest req) // last_id: obj_id obj_type uid
    GetLikedRecordResponse GetPubLikedList(1: GetLikedRecordRequest req) // last_id: obj_id obj_type author
    MGetLikedStateResponse MGetLikedState(1: MGetLikedStateRequest req)
    MGetLikedCountResponse MGetLikedCount(1: MGetLikedCountRequest req)

    // only admin
    EmptyResponse MUpdateLikesCount(1: UpdateLikesCountRequest req)
}
