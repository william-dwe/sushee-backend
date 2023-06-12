package dto

const ACTION_FORGET_PASSWORD = "forget_password"
const ACTION_CHANGE_PASSWORD = "change_password"

const (
	SMTP_CHANGE_PASS_SENDER_NAME = "Blanche"
	SMTP_CHANGE_PASS_SUBJECT     = "Blanche - Change Password Verification Code"
	SMTP_CHANGE_PASS_HTML_PATH   = "template/email/change_password.html"
	SMTP_FORGOT_PASS_SENDER_NAME = "Blanche"
	SMTP_FORGOT_PASS_SUBJECT     = "Blanche - Forgot Password"
	SMTP_FORGOT_PASS_HTML_PATH   = "template/email/forget_password.html"
)

type ChangePasswordRequestVerificationCodeResDTO struct {
	IsEmailSent bool   `json:"is_email_sent"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	RetryIn     int    `json:"retry_in"`
}
type ChangePasswordVerificationCodeReqDTO struct {
	VerificationCode string `json:"verification_code" binding:"required"`
}
type ChangePasswordVerificationCodeResDTO struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

type ForgetPasswordRequestVerificationCodeReqDTO struct {
	Email string `json:"email"`
}
type ForgetPasswordRequestVerificationCodeResDTO struct {
	IsEmailSent bool   `json:"is_email_sent"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	RetryIn     int    `json:"retry_in"`
}
type ForgetPasswordVerificationCodeReqDTO struct {
	VerificationCode string `json:"verification_code" binding:"required"`
}
type ForgetPasswordVerificationCodeResDTO struct {
	AccessToken string `json:"access_token"`
}

type ResetPasswordReqBody struct {
	Password string `json:"password" binding:"required"`
}

type ResetPasswordResBody struct {
	Username string `json:"username"`
}
