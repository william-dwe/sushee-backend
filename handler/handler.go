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
	authUtils        utils.AuthUtil
}

type HandlerConfig struct {
	ExampleUsecase   usecase.ExampleUsecase
	AuthUsecase      usecase.AuthUsecase
	UserUsecase      usecase.UserUsecase
	MenuUsecase      usecase.MenuUsecase
	PromotionUsecase usecase.PromotionUsecase
	CartUsecase      usecase.CartUsecase
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
		authUtils:        c.AuthUtil,
	}
}
