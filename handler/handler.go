package handler

import (
	"sushee-backend/usecase"
)

type Handler struct {
	exampleUsecase usecase.ExampleUsecase
}

type HandlerConfig struct {
	ExampleUsecase usecase.ExampleUsecase
}

func New(c HandlerConfig) *Handler {
	return &Handler{
		exampleUsecase: c.ExampleUsecase,
	}
}
