package app

import (
	"github.com/gin-gonic/gin"
	"kptankhoa.dev/basic-go-gin/configs"
	"kptankhoa.dev/basic-go-gin/internal/database"
)

func InitializeApp() (*gin.Engine, error) {
	configs.LoadEnv(".env")

	dbErr := database.ConnectDatabase()
	if dbErr != nil {
		return nil, dbErr
	}

	router := gin.Default()
	SetUpRoutes(router)

	return router, nil
}
