namespace go user

struct registerRequest {
    1: string username, // Username, up to 32 characters
    2: string password, // Password, up to 32 characters
}

struct registerResponse {
    1: i32 statusCode
    2: string statusMsg
}

struct loginRequest {
    1: string username, // Username, up to 32 characters
    2: string password, // Password, up to 32 characters
}

struct loginResponse {
    1: i32 statusCode,
    2: string statusMsg,
    3: string token,                        // User authentication token
}

service userService {
    registerResponse Register(1: registerRequest req),
    loginResponse Login(1: loginRequest req),
}