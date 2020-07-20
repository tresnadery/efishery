package main

import (
	"os"
	"storage-service/utils"	
	_storageRepository "storage-service/storage/repository/redis"
	_storageUsecase "storage-service/storage/usecase"
	_storageHttpDelivery "storage-service/storage/delivery/http"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main(){	
	rdb := utils.InitRedis()
	r := gin.Default()
	storageRepo := _storageRepository.NewRedisStorageRepository(rdb)
	storageUsecase := _storageUsecase.NewStorageUsecase(storageRepo)
	_storageHttpDelivery.NewStorageHandler(r, storageUsecase)	
	r.Run(":" + os.Getenv("PORT"))	
}