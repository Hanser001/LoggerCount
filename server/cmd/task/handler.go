package main

import (
	"context"
	task "summer/server/shared/kitex_gen/task"
)

// TaskServiceImpl implements the last service interface defined in the IDL.
type TaskServiceImpl struct{}

// NewTask_ implements the TaskServiceImpl interface.
func (s *TaskServiceImpl) NewTask_(ctx context.Context, req *task.NewTaskRequest_) (resp *task.NewTaskResponse_, err error) {
	// TODO: Your code here...
	return
}

// GetTaskStatus implements the TaskServiceImpl interface.
func (s *TaskServiceImpl) GetTaskStatus(ctx context.Context, req *task.GetTaskStatusRequest) (resp *task.GetTaskStatusResponse, err error) {
	// TODO: Your code here...
	return
}
