package app

import (
	"github.com/gin-gonic/gin"
	"kptankhoa.dev/basic-go-gin/internal/database"
	"kptankhoa.dev/basic-go-gin/internal/routes"
	"net/http"
)

func SetUpRoutes(r *gin.Engine) {
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	{
		routes.AlbumRoutes(api)
		routes.UserRoutes(api, database.DB)
	}
}
