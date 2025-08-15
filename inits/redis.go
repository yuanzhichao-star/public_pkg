package inits

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/yuanzhichao-star/public_pkg/config"
	"log"
)

func InitRedis() {
	redisCong := config.AppCong.Redis
	addr := fmt.Sprintf("%s:%d", redisCong.Host, redisCong.Port)
	config.Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisCong.Password, // no password set
		DB:       0,                  // use default DB
	})

	err := config.Rdb.Set(config.Ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	log.Println("redis init success")
}
