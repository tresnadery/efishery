package redis

import(
	"time"
	"context"
	"strconv"
	"github.com/sirupsen/logrus"
	"github.com/go-redis/redis/v8"
	"storage-service/domain"
)

type RedisStorageRepository struct{
	Conn *redis.Client
}

func NewRedisStorageRepository(Conn *redis.Client)*RedisStorageRepository{
	return &RedisStorageRepository{Conn}
}

var ctx = context.Background()

func (r *RedisStorageRepository) Store(rate float64)(error){
	if err := r.Conn.Set(ctx, "rate", rate, 1*time.Hour).Err();err != nil {
    	logrus.Error("failed to store rate : " + err.Error())
        return domain.ErrInternalServerError
    }
    return nil
}

func (r *RedisStorageRepository) GetRate()(float64, error){
	rate, err := r.Conn.Get(ctx, "rate").Result()
	if err == redis.Nil{
		return 0, domain.ErrNotFound
	}
	if err != nil{
		logrus.Error("failed to get rate : " + err.Error())
		return 0, domain.ErrInternalServerError
	}
	convRate, err := strconv.ParseFloat(rate, 64)
	if err != nil{
		logrus.Error("failed to convert rate : " + err.Error())
		return 0, domain.ErrInternalServerError
	}
	return convRate, nil
}