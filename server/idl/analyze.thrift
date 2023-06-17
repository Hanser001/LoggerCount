namespace go analyze

struct analyzeRequest{
    1:string url
    2:i64 user_id
    3:string filed
    4:string obj_name
}

struct analyzeResponse{
    1: i32 statusCode,
    2: string statusMsg,
}

service AnalyzeService{
    analyzeResponse Analyze(1:analyzeRequest req)
}