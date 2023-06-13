package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	RoleName  string         `json:"role_name"`
}
