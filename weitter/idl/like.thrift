namespace go like
include "common.thrift"

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
}

struct GetLikedUsersResponse {
    1: list<i64> uids
}

struct GetLikedRecordRequest {
    1: i64 obj_id
    2: i32 obj_type
    3: i64 uid
    4: i64 author
    5: i64 operator_uid // must equal author or uid
    6: i8 limit
}

struct GetLikedRecordResponse {
    1: list<LikeItem> likes
    2: bool has_more
}

struct GetLikedStateRequest {
    1: i64 obj_id
    2: i32 obj_type
    3: i64 uid
}

struct GetLikedStateResponse {
    1: bool ok
}

struct GetLikedCountRequest {
    1: i64 obj_id
    2: i32 obj_type
}

struct GetLikedCountResponse {
    1: i32 cnt
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

struct UpdateLikeCountRequest {
    1: i64 obj_id
    2: i32 obj_type
    3: i64 uid
    4: i32 count
}

struct MUpdateLikesCountRequest {
    1: map<i32, map<i64, i32>> m
    2: i64 uid
}

service LikeBFF {
    common.EmptyResponse Like(1: LikeRequest req)
    common.EmptyResponse Unlike(1: LikeRequest req)
    GetLikedUsersResponse GetLikedUsers(1: GetLikedUsersRequest req)
    GetLikedRecordResponse GetRcvLikedList(1: GetLikedRecordRequest req)
    GetLikedRecordResponse GetPubLikedList(1: GetLikedRecordRequest req)
    GetLikedStateResponse GetLikedState(1: GetLikedStateRequest req)
    GetLikedCountResponse GetLikedCount(1: GetLikedCountRequest req)
    MGetLikedStateResponse MGetLikedState(1: MGetLikedStateRequest req)
    MGetLikedCountResponse MGetLikedCount(1: MGetLikedCountRequest req)

    // only admin
    common.EmptyResponse UpdateLikeCount(1: UpdateLikeCountRequest req)
    common.EmptyResponse MUpdateLikesCount(1: MUpdateLikesCountRequest req)
}
