package main

import (
	"fmt"
	"kptankhoa.dev/basic-go-gin/routes"
	"net/http"

	"github.com/gin-gonic/gin"

	"kptankhoa.dev/basic-go-gin/configs"
)

func main() {
	configs.LoadEnv(".env")

	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	setUpRoutes(router)

	fmt.Println(configs.PORT)

	err := router.Run(fmt.Sprintf(":%s", configs.PORT))
	if err != nil {
		return
	}

	fmt.Printf("Server is running on port :%s", configs.PORT)
}

func setUpRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		routes.AlbumRoutes(api)
	}
}
