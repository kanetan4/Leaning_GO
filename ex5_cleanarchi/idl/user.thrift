namespace go user

struct User {
    1: i64 id
    2: string name
    3: string email
}

struct CreateUserRequest {
    1: string name
    2: string email
}

struct GetUserRequest {
    1: i64 id
}

service UserService {
    User CreateUser(1: CreateUserRequest request)
    User Getuser(1: GetUserRequest request)
}