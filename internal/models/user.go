package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey" json:"-"`
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid,omitempty"`
	Username string    `gorm:"unique" json:"username,omitempty"`
	Password string    `json:"-"`
}
