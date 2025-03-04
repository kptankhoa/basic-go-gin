package routes

import (
	"github.com/gin-gonic/gin"
	"kptankhoa.dev/basic-go-gin/internal/auth"
	"kptankhoa.dev/basic-go-gin/internal/services"
)

func AlbumRoutes(r *gin.RouterGroup) {
	albumGroup := r.Group("v1/albums")
	albumGroup.Use(auth.Authenticate)
	{
		albumGroup.GET("/", services.GetAlbums)
		albumGroup.GET("/:id", services.GetAlbumByID)
		albumGroup.POST("/albums", services.PostAlbum)
	}
}
