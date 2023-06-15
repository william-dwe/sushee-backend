package handler

import (
	"sushee-backend/usecase"
	"sushee-backend/utils"
)

type Handler struct {
	exampleUsecase   usecase.ExampleUsecase
	authUsecase      usecase.AuthUsecase
	userUsecase      usecase.UserUsecase
	menuUsecase      usecase.MenuUsecase
	promotionUsecase usecase.PromotionUsecase
	cartUsecase      usecase.CartUsecase
	orderUsecase     usecase.OrderUsecase
	paymentUsecase   usecase.PaymentUsecase
	reviewUsecase    usecase.ReviewUsecase
	authUtils        utils.AuthUtil
}

type HandlerConfig struct {
	ExampleUsecase   usecase.ExampleUsecase
	AuthUsecase      usecase.AuthUsecase
	UserUsecase      usecase.UserUsecase
	MenuUsecase      usecase.MenuUsecase
	PromotionUsecase usecase.PromotionUsecase
	CartUsecase      usecase.CartUsecase
	OrderUsecase     usecase.OrderUsecase
	PaymentUsecase   usecase.PaymentUsecase
	ReviewUsecase    usecase.ReviewUsecase
	AuthUtil         utils.AuthUtil
}

func New(c HandlerConfig) *Handler {
	return &Handler{
		exampleUsecase:   c.ExampleUsecase,
		authUsecase:      c.AuthUsecase,
		userUsecase:      c.UserUsecase,
		menuUsecase:      c.MenuUsecase,
		promotionUsecase: c.PromotionUsecase,
		cartUsecase:      c.CartUsecase,
		orderUsecase:     c.OrderUsecase,
		paymentUsecase:   c.PaymentUsecase,
		reviewUsecase:    c.ReviewUsecase,
		authUtils:        c.AuthUtil,
	}
}
