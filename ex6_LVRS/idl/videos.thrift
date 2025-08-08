namespace go videos

struct Video {
    1: i64 id,
    2: string title,
    3: i64 length,
    4: i64 date
    5: string owner,
    6: string status
    7: string format
}

struct VideoList {
    1: list<Video> videos,
    2: i64 total,
}

struct AddVideoRequest {
    1: string title
    2: string format,
    3: i64 length,
    4: i64 date,
    5: string owner,
}

struct GetVideoRequest {
    1: string title
    2: i64 id
}

struct GetVideoListRequest {
    1: i64 page,
    2: i64 pageSize,
    3: string sortBy,
    4: string sortOrder,
    5: string user
}

struct ManualOverrideRequest {
    1: i64 id,
    2: string status,
}

struct DeleteVideoRequest {
    1: i64 id,
    2: string user,
}

service VideoService {
    Video AddVideo(1: AddVideoRequest request)
    bool ManualOverride(1: ManualOverrideRequest request)
    Video GetVideo(1: GetVideoRequest request)
    VideoList GetVideos(1: GetVideoListRequest request)
    bool DeleteVideo(1: DeleteVideoRequest request)
}