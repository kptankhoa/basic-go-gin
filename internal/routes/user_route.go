package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"kptankhoa.dev/basic-go-gin/internal/handlers"
	"kptankhoa.dev/basic-go-gin/internal/repositories"
	"kptankhoa.dev/basic-go-gin/internal/services"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	userGroup := r.Group("v1/user")
	{
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/login", userHandler.Login)
	}
}
