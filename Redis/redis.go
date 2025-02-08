package redis

import (
	"github.com/omniful/go_commons/redis"
)

var Client *redis.Client

func Start() {
	config := &redis.Config{
		Hosts:       []string{"localhost:6379"},
		PoolSize:    50,
		MinIdleConn: 10,
	}

	client := redis.NewClient(config)
	Client = client
}
