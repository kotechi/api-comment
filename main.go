package main

import (
	"net/http"

	config "github.com/ValSpp/gin-gorm-restapi/configs"
	model "github.com/ValSpp/gin-gorm-restapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DatabaseConnection()
	db.Table("users").AutoMigrate(&model.Users{})

	router := gin.Default()
	router.GET("", func(context *gin.Context) {
		context.String(http.StatusOK, "Pong!")
	})

	server := &http.Server{
		Addr:           ":8888",
		Handler:        router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
