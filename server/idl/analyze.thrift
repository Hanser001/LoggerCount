namespace go analyze

struct analyzeRequest{
    1:string url
}

struct analyzeResponse{
      1: i32 statusCode,
      2: string statusMsg,
}