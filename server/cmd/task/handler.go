package main

import (
	"context"
	task "summer/server/shared/kitex_gen/task"
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
	Analyze()
}

// NewTask_ implements the TaskServiceImpl interface.
func (s *TaskServiceImpl) NewTask_(ctx context.Context, req *task.NewTaskRequest_) (resp *task.NewTaskResponse_, err error) {
	resp = new(task.NewTaskResponse_)

	return
}

// GetTaskStatus implements the TaskServiceImpl interface.
func (s *TaskServiceImpl) GetTaskStatus(ctx context.Context, req *task.GetTaskStatusRequest) (resp *task.GetTaskStatusResponse, err error) {
	// TODO: Your code here...
	return
}
