package handler

import (
	"sushee-backend/usecase"
	"sushee-backend/utils"
)

type Handler struct {
	exampleUsecase usecase.ExampleUsecase
	authUsecase    usecase.AuthUsecase
	userUsecase    usecase.UserUsecase
	authUtils      utils.AuthUtil
}

type HandlerConfig struct {
	ExampleUsecase usecase.ExampleUsecase
	AuthUsecase    usecase.AuthUsecase
	UserUsecase    usecase.UserUsecase
	AuthUtil       utils.AuthUtil
}

func New(c HandlerConfig) *Handler {
	return &Handler{
		exampleUsecase: c.ExampleUsecase,
		authUsecase:    c.AuthUsecase,
		userUsecase:    c.UserUsecase,
		authUtils:      c.AuthUtil,
	}
}
