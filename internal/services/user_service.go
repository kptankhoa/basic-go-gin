package services

import (
	"errors"
	"gorm.io/gorm"
	"kptankhoa.dev/basic-go-gin/internal/auth"
	"kptankhoa.dev/basic-go-gin/internal/models"
	"kptankhoa.dev/basic-go-gin/internal/repositories"
)

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserService struct {
	userRepository *repositories.UserRepository
}

var ErrUserNotFound = errors.New("user not found")
var ErrUserExisted = errors.New("username already existed")

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) RegisterUser(input RegisterInput) error {
	_, err := s.userRepository.GetUserByUsername(input.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrUserExisted
	}
	hashedPassword := auth.HashPassword(input.Password)
	user := models.User{Username: input.Username, Password: hashedPassword}
	return s.userRepository.CreateUser(user)
}

func (s *UserService) LoginUser(input LoginInput) (string, error) {
	user, err := s.userRepository.GetUserByUsername(input.Username)
	if err != nil || auth.HashPassword(input.Password) != user.Password {
		return "", ErrUserNotFound
	}

	return auth.GenerateToken(user.Username)
}
