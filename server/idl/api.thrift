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
    1: string username(api.query="username", api.vd="len($)>0 && len($)<33"), // Username, up to 32 characters
    2: string password(api.query="password", api.vd="len($)>0 && len($)<33"), // Password, up to 32 characters
}

struct loginResponse {
    1: i32 statusCode,
    2: string statusMsg,
    3: string token,                        // User authentication token
}

struct uploadFileRequest{
    1: string token(api.query="token") // User authentication token
    2: string file_path(api.query="path")
}

struct uploadFileResponse{
     1: i32 statusCode,
     2: string statusMsg,
}

struct downloadFileRequest{
    1: string token(api.query="token"),
    2: string file_name(api.query="file_name"),
}

struct downloadFileResponse{
     1: i32 statusCode,
     2: string statusMsg,
}

struct newTaskRequest{
    1: string token(api.query="token"),

}

struct newTaskResponse{
     1: i32 statusCode,
     2: string statusMsg,
}

struct getTaskRequest{
    1: string token(api.query="token"),
    2: string task_id(api.query="task_id"),
}

struct getTaskResponse{
    1: i32 statusCode,
    2: string statusMsg,
    3: string task_id,
    4: i32 code,  // 0 success,1 working or failed
}

service ApiService {
    registerResponse Register(1:registerRequest req)(api.post="/api/user/register/")
    loginResponse Login(1:loginRequest req)(api.post="/api/user/login/")

    uploadFileResponse Upload(1: uploadFileRequest req)(api.post="api/file")
    downloadFileResponse Download(1: downloadFileRequest req)(api.get="api/file")

    newTaskRequest NewTask(1: newTaskRequest req)(api.post="api/task")
    getTaskResponse GetTask(1: getTaskRequest req)(api.get="api/task")
}