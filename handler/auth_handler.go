package handler

import (
	"sushee-backend/config"
	"sushee-backend/httperror/domain"
	"sushee-backend/utils"

	"sushee-backend/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var reqBody dto.UserRegisterReqBody
	if err := utils.ShouldBindJsonWithValidation(c, &reqBody); err != nil {
		_ = c.Error(domain.ErrAuthHandlerRegisterIncompleteReqBody)
		return
	}

	u, err := h.authUsecase.Register(&reqBody)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_POST_REGISTER_USER",
		Message: "Success post register user",
		Data:    u,
	}
	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) Login(c *gin.Context) {
	var reqBody dto.UserLoginReqBody
	if err := utils.ShouldBindJsonWithValidation(c, &reqBody); err != nil {
		_ = c.Error(domain.ErrAuthHandlerLoginIncompleteReqBody)
		return
	}

	l, accessToken, refreshToken, err := h.authUsecase.Login(&reqBody)
	if err != nil {
		_ = c.Error(err)
		return
	}

	isUserLoggedIn := true
	h.authUtils.EmbedTokenOnContextCookie(
		c,
		&refreshToken,
		&accessToken,
		&isUserLoggedIn,
		config.Config.AppConfig.Url,
	)

	res := dto.ResponseStruct{
		Code:    "SUCCESS_POST_LOGIN_USER",
		Message: "Success post login user",
		Data:    l,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		_ = c.Error(domain.ErrAuthHandlerRefreshTokenNotFound)
		return
	}

	l, err := h.authUsecase.Logout(refreshToken)
	if err != nil {
		_ = c.Error(err)
		return
	}

	newAccessToken := ""
	newRefreshToken := ""
	isUserLoggedIn := false
	h.authUtils.EmbedTokenOnContextCookie(
		c,
		&newRefreshToken,
		&newAccessToken,
		&isUserLoggedIn,
		config.Config.AppConfig.Url,
	)

	res := dto.ResponseStruct{
		Code:    "SUCCESS_POST_LOGOUT_USER",
		Message: "Success post logout user",
		Data:    l,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		_ = c.Error(domain.ErrAuthHandlerRefreshTokenNotFound)
		return
	}

	r, accessToken, err := h.authUsecase.Refresh(refreshToken)
	if err != nil {
		_ = c.Error(err)
		return
	}

	isUserLoggedIn := true
	h.authUtils.EmbedTokenOnContextCookie(
		c,
		nil,
		&accessToken,
		&isUserLoggedIn,
		config.Config.AppConfig.Url,
	)

	res := dto.ResponseStruct{
		Code:    "SUCCESS_POST_REFRESH_USER",
		Message: "Success post refresh user",
		Data:    r,
	}

	utils.ResponseSuccessJSON(c, res)
}

// func (h *Handler) AuthenticateRole(c *gin.Context) {
// 	username := c.GetString("username")
// 	roleFromToken := c.GetString("role")

// 	user, err := h.authUsecase.GetDetailUserByUsername(username)
// 	if err != nil {
// 		router_helper.GenerateErrorMessage(c, err)
// 		return
// 	}

// 	roleFromDb, err := h.authUsecase.GetDetailRole(user.RoleId)
// 	if err != nil {
// 		router_helper.GenerateErrorMessage(c, err)
// 		return
// 	}

// 	if roleFromToken != roleFromDb.RoleName {
// 		router_helper.GenerateErrorMessage(c, httperror.UnauthorizedError())
// 		return
// 	}

// 	c.Next()
// }

// func (h *Handler) UserAuthorization(c *gin.Context) {
// 	role := c.GetString("role")
// 	if role != "user" {
// 		router_helper.GenerateErrorMessage(c, httperror.ForbiddenError())
// 		return
// 	}
// 	c.Next()
// }

// func (h *Handler) AdminAuthorization(c *gin.Context) {
// 	role := c.GetString("role")
// 	if role != "admin" {
// 		router_helper.GenerateErrorMessage(c, httperror.ForbiddenError())
// 		return
// 	}
// 	c.Next()
// }
