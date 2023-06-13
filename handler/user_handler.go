package handler

import (
	"sushee-backend/dto"
	"sushee-backend/httperror/domain"
	"sushee-backend/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShowUserDetail(c *gin.Context) {
	username := c.GetString("username")

	userInfo, err := h.userUsecase.GetDetailUserByUsername(username)
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
	username := c.GetString("username")
	var formFile dto.UserProfileUploadReqBody

	if err := c.ShouldBind(&formFile); err != nil {
		_ = c.Error(domain.ErrUserHandlerUpdateIncompleteReqBody)
		return
	}

	updateProfileReq := dto.UserEditDetailsReqBody{
		FullName: formFile.FullName,
		Phone:    formFile.Phone,
		Email:    formFile.Email,
		Password: formFile.Password,
	}

	userInfo, err := h.userUsecase.UpdateUserDetailsByUsername(username, updateProfileReq)
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
