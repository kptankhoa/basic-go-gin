package auth

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"-"`
	UUID      uuid.UUID      `gorm:"uniqueIndex,type:uuid;default:uuid_generate_v4()" json:"uuid,omitempty"`
	Username  string         `gorm:"uniqueIndex" json:"username,omitempty"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
