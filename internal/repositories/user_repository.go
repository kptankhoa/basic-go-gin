package repositories

import (
	"gorm.io/gorm"
	"kptankhoa.dev/basic-go-gin/internal/auth"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user auth.User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) GetUserByUsername(username string) (auth.User, error) {
	var user auth.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}
