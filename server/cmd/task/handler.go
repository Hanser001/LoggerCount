package main

import (
	"context"
	"summer/server/shared/errno"
	task "summer/server/shared/kitex_gen/task"
	"time"
)

// TaskServiceImpl implements the last service interface defined in the IDL.
type TaskServiceImpl struct {
	RedisManger
	AnalyzeManger
}

type RedisManger interface {
	SetTaskRecord(ctx context.Context, userId int, taskId int) error
	UpdateTaskStatus(ctx context.Context, userId int, taskId int) error
}

type AnalyzeManger interface {
	Analyze(ctx context.Context, url, filed, objname string, userId int64)
}

// NewTask_ implements the TaskServiceImpl interface.
func (s *TaskServiceImpl) NewTask_(ctx context.Context, req *task.NewTaskRequest_) (resp *task.NewTaskResponse_, err error) {
	resp = new(task.NewTaskResponse_)

	//ignore logic that creates new task
	//after task was created,run redis and analyze server
	taskId := time.Now().Unix()
	s.RedisManger.SetTaskRecord(ctx, int(req.UserId), int(taskId))

	go s.AnalyzeManger.Analyze(ctx, "example1", "example2", "example3", req.UserId)

	resp.StatusCode = int32(errno.Success.ErrCode)
	resp.StatusMsg = errno.Success.ErrMsg

	return
}

// GetTaskStatus implements the TaskServiceImpl interface.
func (s *TaskServiceImpl) GetTaskStatus(ctx context.Context, req *task.GetTaskStatusRequest) (resp *task.GetTaskStatusResponse, err error) {
	// TODO: Your code here...
	return
}
