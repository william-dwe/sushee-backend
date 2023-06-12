package handler

import (
	"sushee-backend/httperror/domain"

	"sushee-backend/dto"
	"sushee-backend/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ExampleHandler(c *gin.Context) {
	var inputRequest dto.ExampleReqDTO
	if err := utils.ShouldBindJsonWithValidation(c, &inputRequest); err != nil {
		utils.ResponseErrorJSON(c, err)
		return
	}

	res, err := h.exampleUsecase.ExampleProcess(inputRequest)
	if err != nil {
		utils.ResponseErrorJSON(c, err)
		return
	}

	response := utils.ResponseStruct{
		Code:    "SUCCEED_EXAMPLE_HANDLER",
		Message: "Success",
		Data:    res,
	}

	utils.ResponseSuccessJSON(c, response)
}

func (h *Handler) ExampleHandlerErrorMiddleware(c *gin.Context) {
	_ = c.Error(domain.ErrExampleUnexpected)
}
