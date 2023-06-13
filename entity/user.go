package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName       string
	Phone          string
	Email          string
	Username       string
	Password       string
	RegisterDate   time.Time
	ProfilePicture *string
	PlayAttempt    int
	RoleId         int
}
