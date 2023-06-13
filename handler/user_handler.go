package handler

import (
	"sushee-backend/dto"
	"sushee-backend/httperror/domain"
	"sushee-backend/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShowUserDetail(c *gin.Context) {
	user, err := utils.GetUserJWTContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	userInfo, err := h.userUsecase.GetDetailUserByUsername(user.Username)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_SHOW_USER_DETAIL",
		Message: "Success get user detail",
		Data:    userInfo,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) UpdateUserProfile(c *gin.Context) {
	user, err := utils.GetUserJWTContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	var reqBody dto.UserProfileReqBody

	if err := utils.ShouldBindWithValidation(c, &reqBody); err != nil {
		_ = c.Error(domain.ErrUserHandlerUpdateIncompleteReqBody)
		return
	}

	userInfo, err := h.userUsecase.UpdateUserDetailsByUsername(user.Username, reqBody)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_UPDATE_USER_PROFILE",
		Message: "Success update user profile",
		Data:    userInfo,
	}

	utils.ResponseSuccessJSON(c, res)
}
