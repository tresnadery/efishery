package main

import (
	"os"
	"time"	
	"auth-service/utils"	

	_userRepo "auth-service/user/repository/pgsql"
	_userUsecase "auth-service/user/usecase"
	_userHttpDelivery "auth-service/user/delivery/http"

	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main(){
	dbConn := utils.GetDBConnection()
	defer dbConn.Close()
	r := gin.Default()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	userRepo := _userRepo.NewPgsqlUserRepository(dbConn)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)
	_userHttpDelivery.NewUserHandler(r, userUsecase, dbConn)

	r.Run(":" + os.Getenv("PORT"))	
}