package dto

import "mime/multipart"

type UserProfileReqBody struct {
	FullName       string                `form:"full_name,omitempty"`
	Email          string                `form:"email,omitempty"`
	Phone          string                `form:"phone,omitempty"`
	Password       string                `form:"password,omitempty"`
	ProfilePicture *multipart.FileHeader `form:"profile_picture,omitempty"`
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
