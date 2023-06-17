namespace go task

struct newTaskRequest{
    1: i64 user_id
}

struct newTaskResponse{
     1: i32 statusCode,
     2: string statusMsg,
}

struct TaskStatus{
    1: string task_id,
    2: i32 status_code,
}

struct GetTaskStatusRequest{
    1: i64 user_id,
}

struct GetTaskStatusResponse{
     1: i32 statusCode,
     2: string statusMsg,
     3: TaskStatus serverResponse,
}

service TaskService{
    newTaskResponse NewTask(newTaskRequest req),
    GetTaskStatusResponse GetTaskStatus(GetTaskStatusRequest req),
}