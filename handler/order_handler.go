package handler

import (
	"strconv"
	"sushee-backend/dto"
	"sushee-backend/httperror/domain"
	"sushee-backend/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPaymentOption(c *gin.Context) {
	paymentOptions, err := h.paymentUsecase.GetPaymentOption()
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_GET_PAYMENT_OPTION",
		Message: "Success get payment option",
		Data:    paymentOptions,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) AddOrder(c *gin.Context) {
	user, err := utils.GetUserJWTContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	var reqBody dto.OrderReqBody
	if err := utils.ShouldBindWithValidation(c, &reqBody); err != nil {
		_ = c.Error(err)
		return
	}

	order, err := h.orderUsecase.AddOrder(user.Username, &reqBody)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_ADD_ORDER",
		Message: "Success add order",
		Data:    order,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) DeleteOrder(c *gin.Context) {
	user, err := utils.GetUserJWTContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	var reqBody dto.OrderReqBody
	if err := utils.ShouldBindWithValidation(c, &reqBody); err != nil {
		_ = c.Error(err)
		return
	}

	order, err := h.orderUsecase.AddOrder(user.Username, &reqBody)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_DELETE_ORDER",
		Message: "Success delete order",
		Data:    order,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) GetOrders(c *gin.Context) {
	user, err := utils.GetUserJWTContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil {
		_ = c.Error(domain.ErrOrderHandlerIncorrectLimitRequest)
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		_ = c.Error(domain.ErrOrderHandlerIncorrectPageRequest)
		return
	}
	q := dto.OrderHistoryQuery{
		Search:           c.DefaultQuery("s", "%"),
		SortBy:           c.DefaultQuery("sortBy", "id"),
		FilterByCategory: c.DefaultQuery("filterByCategory", ""),
		Sort:             c.DefaultQuery("sort", "desc"),
		Limit:            limit,
		Page:             page,
	}

	orders, err := h.orderUsecase.GetOrderHistory(user.Username, &q)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_GET_ORDER",
		Message: "Success get order",
		Data:    orders,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) GetOrderStatus(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil {
		_ = c.Error(domain.ErrOrderHandlerIncorrectLimitRequest)
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		_ = c.Error(domain.ErrOrderHandlerIncorrectPageRequest)
		return
	}
	q := dto.OrderStatusQuery{
		Search:         c.DefaultQuery("s", "%"),
		SortBy:         c.DefaultQuery("sortBy", "id"),
		FilterByStatus: c.DefaultQuery("filterByStatus", ""),
		Sort:           c.DefaultQuery("sort", "desc"),
		Limit:          limit,
		Page:           page,
	}

	orders, err := h.orderUsecase.GetOrderStatus(&q)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_GET_ORDER_STATUS",
		Message: "Success get order status",
		Data:    orders,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	var reqBody dto.OrderStatusUpdateReqBody
	if err := utils.ShouldBindWithValidation(c, &reqBody); err != nil {
		_ = c.Error(domain.ErrUserHandlerUpdateIncompleteReqBody)
		return
	}

	orders, err := h.orderUsecase.UpdateOrderStatus(&reqBody)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_UPDATE_ORDER_STATUS",
		Message: "Success update order status",
		Data:    orders,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) AddReview(c *gin.Context) {
	user, err := utils.GetUserJWTContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	var reqBody dto.ReviewAddReqBody
	if err := utils.ShouldBindWithValidation(c, &reqBody); err != nil {
		_ = c.Error(domain.ErrUserHandlerUpdateIncompleteReqBody)
		return
	}

	order, err := h.reviewUsecase.AddReview(user.Username, &reqBody)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_ADD_REVIEW",
		Message: "Success add review",
		Data:    order,
	}

	utils.ResponseSuccessJSON(c, res)
}
