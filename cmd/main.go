package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kptankhoa.dev/basic-go-gin/internal/database"
	"kptankhoa.dev/basic-go-gin/internal/routes"
	"net/http"

	"kptankhoa.dev/basic-go-gin/configs"
)

func main() {
	configs.LoadEnv(".env")

	database.ConnectDatabase()

	router := gin.Default()
	setUpRoutes(router)

	err := router.Run(fmt.Sprintf(":%s", configs.PORT))
	if err != nil {
		return
	}

	fmt.Printf("Server is running on port :%s", configs.PORT)
}

func setUpRoutes(r *gin.Engine) {
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		routes.AlbumRoutes(api)
	}
}
