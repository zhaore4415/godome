package cache

import (
	"context"
	"crypto/tls"
	"time"

	"myproject/internal/proto/hello"

	redis "github.com/redis/go-redis/v9"
)

const (
	prefix = "bsi:hello"
)

func NewRedisCache(bs *hello.Bootstrap) (*redis.Client, error) {
	kv := redis.NewClient(&redis.Options{
		Addr:         bs.Data.Redis.Addr,
		Password:     bs.Data.Redis.Pwd,
		DB:           int(bs.Data.Redis.Db),
		PoolSize:     30,
		MaxRetries:   3,
		ReadTimeout:  time.Duration(bs.Data.Redis.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(bs.Data.Redis.WriteTimeout) * time.Second,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})
	return kv, kv.Ping(context.TODO()).Err()
}
