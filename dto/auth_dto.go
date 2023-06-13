package dto

import (
	"mime/multipart"
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

type UserEditDetailsReqBody struct {
	FullName       string `json:"full_name"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profile_picture"`
}

type UserProfileUploadReqBody struct {
	FullName string                `form:"full_name,omitempty"`
	Email    string                `form:"email,omitempty"`
	Phone    string                `form:"phone,omitempty"`
	Password string                `form:"password,omitempty"`
	Img      *multipart.FileHeader `form:"img,omitempty"`
}

type UserContext struct {
	Username       string `json:"username"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
	PlayAttempt    int    `json:"play_attempt"`
	RoleId         int    `json:"role_id"`
}
