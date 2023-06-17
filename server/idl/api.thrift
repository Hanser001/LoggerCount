namespace go api

struct registerRequest {
    1: string username(api.query="username", api.vd="len($)>0 && len($)<33") // Username, up to 32 characters
    2: string password(api.query="password", api.vd="len($)>0 && len($)<33") // Password, up to 32 characters
}

struct registerResponse {
    1: i32 statusCode
    2: string statusMsg
    3: string token
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

service ApiService {
    registerResponse Register(1:registerRequest req)(api.post="/api/user/register/")
    loginResponse Login(1:loginRequest req)(api.post="/api/user/login/")
}