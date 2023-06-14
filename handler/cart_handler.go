package handler

import (
	"strconv"
	"sushee-backend/dto"
	"sushee-backend/httperror/domain"
	"sushee-backend/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddCart(c *gin.Context) {
	user, err := utils.GetUserJWTContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var reqBody dto.CartReqBody
	if err := utils.ShouldBindWithValidation(c, &reqBody); err != nil {
		_ = c.Error(domain.ErrCartHandlerIncompleteReqBody)
		return
	}

	cart, err := h.cartUsecase.AddCart(user.Username, &reqBody)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_ADD_CART_ITEM",
		Message: "Success add cart item",
		Data:    cart,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) ShowCart(c *gin.Context) {
	user, err := utils.GetUserJWTContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	t, err := h.cartUsecase.GetCart(user.Username)
	if err != nil {
		_ = c.Error(err)
		return
	}

	resBody := dto.CartsResBody{
		Carts: *t,
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_SHOW_CART",
		Message: "Success get cart items",
		Data:    resBody,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) DeleteCarts(c *gin.Context) {
	user, err := utils.GetUserJWTContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = h.cartUsecase.DeleteCartsByUsername(user.Username)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_DELETE_CART",
		Message: "Success delete cart items",
		Data:    nil,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) DeleteCartById(c *gin.Context) {
	user, err := utils.GetUserJWTContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		_ = c.Error(domain.ErrCartHandlerInvalidCartId)
		return
	}

	err = h.cartUsecase.DeleteCartByCartId(user.Username, cartId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_DELETE_CART",
		Message: "Success delete cart items",
		Data:    nil,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) UpdateCartById(c *gin.Context) {
	user, err := utils.GetUserJWTContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	cartId, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		_ = c.Error(domain.ErrCartHandlerInvalidCartId)
		return
	}
	var reqBody dto.CartEditDetailsReqBody
	if err := utils.ShouldBindWithValidation(c, &reqBody); err != nil {
		_ = c.Error(domain.ErrCartHandlerIncompleteReqBody)
		return
	}

	data, err := h.cartUsecase.UpdateCartByCartId(user.Username, cartId, &reqBody)
	if err != nil {
		_ = c.Error(err)
		return
	}
	res := dto.ResponseStruct{
		Code:    "SUCCESS_UPDATE_CART",
		Message: "Success update cart items",
		Data:    data,
	}

	utils.ResponseSuccessJSON(c, res)

}
