package routes

import (
	"github.com/gin-gonic/gin"
	"kptankhoa.dev/basic-go-gin/controllers"
)

func AlbumRoutes(r *gin.RouterGroup) {
	albumGroup := r.Group("v1/albums")
	{
		albumGroup.GET("/", controllers.GetAlbums)
		albumGroup.GET("/:id", controllers.GetAlbumByID)
		albumGroup.POST("/albums", controllers.PostAlbum)
	}
}
