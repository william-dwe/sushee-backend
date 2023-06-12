package server

import (
	"fmt"

	"sushee-backend/db"
	"sushee-backend/repository"
	"sushee-backend/usecase"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	exampleRepo := repository.NewExampleRepository(repository.ExampleRepositoryConfig{
		DB: db.Get(),
	})

	exampleUsecase := usecase.NewExampleUsecase(usecase.ExampleUsecaseConfig{
		ExampleRepository: exampleRepo,
	})

	r := CreateRouter(RouterConfig{
		ExampleUsecase: exampleUsecase,
	})
	return r
}

func Init() {
	r := initRouter()
	err := r.Run()
	if err != nil {
		fmt.Println("error while running server", err)
		return
	}
}
