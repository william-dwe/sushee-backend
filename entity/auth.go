package entity

import (
	"time"

	"gorm.io/gorm"
)

type AuthTokenPayload struct {
	Username string `json:"username"`
}

type AuthSession struct {
	ID           uint `gorm:"primaryKey"`
	UserId       uint
	RefreshToken string
	IsInvalid    bool
	ExpiredAt    time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
