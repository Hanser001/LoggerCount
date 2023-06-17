namespace go file

struct uploadRequest {
    1: i64 user_id,      // User Id
    2: string file_path,
}

struct uploadResponse {
     1: i32 statusCode,
     2: string statusMsg,
     3: string url,
}

struct downloadRequest{
    1: i64 user_id,
    2: string objectName,
}

struct downloadResponse{
     1: i32 statusCode,
     2: string statusMsg,
     3: string url,
}

service fileService{
    uploadResponse UploadFile(1:uploadRequest req),
    downloadResponse DownloadFile(1:downloadRequest req)
}