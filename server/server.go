package server

import (
	"fmt"

	"sushee-backend/db"
	"sushee-backend/repository"
	"sushee-backend/usecase"
	"sushee-backend/utils"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	authUtil := utils.NewAuthUtil()
	clientUploader := utils.NewClientUploader()
	gcsUtil := utils.NewGCSUploader(utils.GCSUploaderConfig{
		ClientUploader: clientUploader,
	})

	exampleRepo := repository.NewExampleRepository(repository.ExampleRepositoryConfig{
		DB: db.Get(),
	})
	userRepo := repository.NewUserRepository(repository.UserRepositoryConfig{
		DB: db.Get(),
	})
	authRepo := repository.NewAuthRepository(repository.AuthRepositoryConfig{
		DB: db.Get(),
	})

	exampleUsecase := usecase.NewExampleUsecase(usecase.ExampleUsecaseConfig{
		ExampleRepository: exampleRepo,
	})
	mediaUsecase := usecase.NewMediaUsecase(usecase.MediaUsecaseConfig{
		GCSUploader: gcsUtil,
	})
	userUsecase := usecase.NewUserUsecase(usecase.UserUsecaseConfig{
		UserRepository: userRepo,
		MediaUsecase:   mediaUsecase,
	})

	authUsecase := usecase.NewAuthUsecase(usecase.AuthUsecaseConfig{
		AuthRepository: authRepo,
		UserRepository: userRepo,
		AuthUtil:       authUtil,
	})

	r := CreateRouter(RouterConfig{
		ExampleUsecase: exampleUsecase,
		UserUsecase:    userUsecase,
		AuthUsecase:    authUsecase,
		AuthUtil:       authUtil,
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
