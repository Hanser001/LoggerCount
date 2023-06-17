package pkg

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type RedisManger struct {
	client *redis.Client
}

// NewRedisManager creates a new redis manager.
func NewRedisManager(client *redis.Client) *RedisManger {
	return &RedisManger{client: client}
}

func (r *RedisManger) NewUpload(ctx context.Context, userId int64, filed string, value int64) error {
	//TODO: add error check
	r.client.ZAdd(ctx, strconv.FormatInt(userId, 10), &redis.Z{
		Score:  float64(value),
		Member: filed,
	})
	return nil
}
