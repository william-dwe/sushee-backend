package server

import (
	"sushee-backend/handler"
	"sushee-backend/httperror"
	"sushee-backend/middleware"
	"sushee-backend/usecase"
	"sushee-backend/utils"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	ExampleUsecase usecase.ExampleUsecase
	UserUsecase    usecase.UserUsecase
	AuthUsecase    usecase.AuthUsecase
	MenuUsecase    usecase.MenuUsecase
	AuthUtil       utils.AuthUtil
}

func CreateRouter(c RouterConfig) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorHandler)

	h := handler.New(handler.HandlerConfig{
		ExampleUsecase: c.ExampleUsecase,
		UserUsecase:    c.UserUsecase,
		AuthUsecase:    c.AuthUsecase,
		MenuUsecase:    c.MenuUsecase,
		AuthUtil:       c.AuthUtil,
	})

	r.NoRoute(func(c *gin.Context) {
		utils.ResponseErrorJSON(c, httperror.NotFoundError("endpoint not found"))
	})
	r.GET("/ping", func(c *gin.Context) {
		utils.ResponseSuccessJSONData(c, "pong!")
	})

	apiEndpoint := r.Group("/api")
	v1 := apiEndpoint.Group("/v1")
	selectedVersion := v1.Group("")

	selectedVersion.POST("/example-process", h.ExampleHandler)
	selectedVersion.POST("/example-process-error", h.ExampleHandlerErrorMiddleware)

	selectedVersion.POST("/login", h.Login)
	selectedVersion.POST("/register", h.Register)
	selectedVersion.GET("/menus", h.ShowMenu)
	selectedVersion.GET("/promotions", h.ShowPromotion)

	a := selectedVersion.Group("/")
	a.Use(middleware.Authenticate)

	a.POST("/logout", h.Logout)
	a.POST("/refresh", h.Refresh)

	user := a.Group("/users")
	user.GET("/me", h.ShowUserDetail)
	user.POST("/me", h.UpdateUserProfile)

	return r
}
