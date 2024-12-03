package redisConn

import (
	"context"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/redis/go-redis/v9"
)

func RedisConn(pctx context.Context, cfg *config.Config) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.RedisUrl,
		Password: "",
		DB:       0,
	})

	return client

}
