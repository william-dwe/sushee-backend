package dto

import (
	"time"
)

type UserLoginReqBody struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
type UserLoginResBody struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
	Email       string `json:"email"`
}

type UserRegisterReqBody struct {
	FullName string `json:"full_name" binding:"required"`
	Phone    string `json:"phone"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserRegisterResBody struct {
	FullName       string    `json:"full_name"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	RegisterDate   time.Time `json:"register_date"`
	ProfilePicture string    `json:"profile_picture"`
	RoleId         int       `json:"role_id"`
}

type UserLogoutResBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
