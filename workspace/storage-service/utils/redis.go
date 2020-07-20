package utils

import(
	"os"
	"log"
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedis() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
        Password: "", // no password set
        DB:       0,  // use default DB
    })    
    pong, err := rdb.Ping(ctx).Result()
    log.Println(pong, err)
    return rdb
}