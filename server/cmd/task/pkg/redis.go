package pkg

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"summer/server/shared/consts"
)

type RedisManger struct {
	client *redis.Client
}

// NewRedisManager creates a new redis manager.
func NewRedisManager(client *redis.Client) *RedisManger {
	return &RedisManger{client: client}
}

func (m *RedisManger) SetTaskRecord(ctx context.Context, userId int, taskId int) error {
	// TODO: add error check
	m.client.HSet(ctx, strconv.Itoa(userId),
		strconv.Itoa(taskId),
		consts.TaskWorking,
	)
	return nil
}

func (m *RedisManger) UpdateTaskStatus(ctx context.Context, userId int, taskId int) error {
	// TODO: add error check
	m.client.HSet(ctx, strconv.Itoa(userId),
		strconv.Itoa(taskId),
		consts.TaskSuccess,
	)
	return nil
}
