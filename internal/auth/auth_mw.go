package auth

import (
	"github.com/gin-gonic/gin"
	"kptankhoa.dev/basic-go-gin/internal/utils"
	"net/http"
)

func Authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		c.Abort()
		return
	}

	username, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("username", username)
	c.Next()
}
