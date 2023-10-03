package messaging

import (
	"github.com/abc-valera/flugo-api-golang/internal/domain/service"
	"github.com/hibiken/asynq"
)

type redisTaskDistributor struct {
	client *asynq.Client

	log service.Logger
}

func newRedisTaskDistributor(redisOpt *asynq.RedisClientOpt, log service.Logger) *redisTaskDistributor {
	return &redisTaskDistributor{
		log:    log,
		client: asynq.NewClient(redisOpt),
	}
}
