package entity

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName  string
	Phone string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Username string `gorm:"unique"`
	Password  string 
	RegisterDate time.Time
	ProfilePicture string
	PlayAttempt int
	RoleId int
}

type UserLoginReqBody struct {
	Identifier    string `json:"identifier" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResBody struct {
	AccessToken string `json:"access_token"`
	RoleName string `json:"role_name"`
}

type UserRegisterReqBody struct {
	FullName  string `json:"full_name" binding:"required"`
	Phone string `gorm:"unique"`
	Email     string `gorm:"unique" binding:"required"`
	Username string `gorm:"unique" binding:"required"`
	Password  string `binding:"required"`
}

type UserRegisterResBody struct {
	FullName  string
	Phone string
	Email     string
	Username string
	Password  string 
	RegisterDate time.Time
	ProfilePicture string
	PlayAttempt int
	RoleId int
}

type UserContext struct {
	Username string `json:"username"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Phone string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
	PlayAttempt int `json:"play_attempt"`
	RoleId int `json:"role_id"`
}

type UserEditDetailsReqBody struct {
	FullName  string 
	Phone 	string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	ProfilePicture string 
}

type UserProfileUploadReqBody struct {
	FullName string `form:"full_name,omitempty"`
	Email string `form:"email,omitempty"`
	Phone string `form:"phone,omitempty"`
	Password  string `form:"password,omitempty"`
	Img *multipart.FileHeader `form:"img,omitempty"`
}
