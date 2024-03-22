package utils

import (
	"fmt"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"log"
)

type Redis struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func InitRedisRedsync(r *Redis) *redsync.Redsync {
	RedisClient := goredislib.NewClient(&goredislib.Options{
		Addr: fmt.Sprintf("%s:%d", r.Host, r.Port),
	})
	pool := goredis.NewPool(RedisClient)
	rs := redsync.New(pool)
	return rs
}
func MutexUnlock(id int, rs *redsync.Redsync) {
	mutex := rs.NewMutex(fmt.Sprintf("goods_%d", id))
	if err := mutex.Lock(); err != nil {
		log.Fatal("获取redis分布式锁失败")
	}
	defer func() {
		if _, err := mutex.Unlock(); err != nil {
			log.Printf("释放redis分布式锁失败: %v", err)
		}
	}()
}
