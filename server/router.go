package server

import (
	"sushee-backend/handler"
	"sushee-backend/httperror"
	"sushee-backend/usecase"
	"sushee-backend/utils"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	ExampleUsecase usecase.ExampleUsecase
}

func CreateRouter(c RouterConfig) *gin.Engine {
	r := gin.Default()
	// r.Use(middleware.JSONifyResult())

	h := handler.New(handler.HandlerConfig{
		ExampleUsecase: c.ExampleUsecase,
	})

	r.NoRoute(func(c *gin.Context) {
		utils.ResponseErrorJSON(c, httperror.NotFoundError("endpoint not found"))
	})
	r.GET("/ping", func(c *gin.Context) {
		utils.ResponseSuccessJSONData(c, "pong!")
	})

	apiEndpoint := r.Group("/api")
	v1 := apiEndpoint.Group("/v1")

	v1.POST("/example-process", h.ExampleHandler)
	v1.POST("/example-process-error", h.ExampleHandlerErrorMiddleware)

	return r
}
