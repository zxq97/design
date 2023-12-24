namespace go common

struct EmptyResponse {
}

struct CheckResponse {
    1: bool ok
}

struct GetUserItemRequest {
    1: i64 uid
    2: i64 last_id
    3: i8 limit
}

struct GetItemRequest {
    1: i64 id
}

struct MGetItemARequest {
    1: list<i64> ids
}
