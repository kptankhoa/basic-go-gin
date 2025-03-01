package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"kptankhoa.dev/basic-go-gin/internal/services"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var input services.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.userService.RegisterUser(input)
	if errors.Is(err, services.ErrUserExisted) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to register", "reason": fmt.Sprint(err)})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register", "reason": fmt.Sprint(err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var input services.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := h.userService.LoginUser(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
