package utils

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"log"
)

type Redis struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

var RedisClient *redis.Client

func InitRedis(r *Redis) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", r.Host, r.Port),
	})
}
func InitRedisRedsync() *redsync.Redsync {
	pool := goredis.NewPool(RedisClient)
	rs := redsync.New(pool)
	return rs
}
func MutexUnlock(id int, rs *redsync.Redsync) *redsync.Mutex {
	mutex := rs.NewMutex(fmt.Sprintf("goods_%d", id))
	if err := mutex.Lock(); err != nil {
		log.Fatal("获取redis分布式锁失败")
	}
	return mutex
}
