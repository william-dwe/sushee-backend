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
	menuRepo := repository.NewMenuRepository(repository.MenuRepositoryConfig{
		DB: db.Get(),
	})
	promotionRepo := repository.NewPromotionRepository(repository.PromotionRepositoryConfig{
		DB: db.Get(),
	})
	cartRepo := repository.NewCartRepository(repository.CartRepositoryConfig{
		DB: db.Get(),
	})
	couponRepo := repository.NewCouponRepository(repository.CouponRepositoryConfig{
		DB: db.Get(),
	})
	orderRepo := repository.NewOrderRepository(repository.OrderRepositoryConfig{
		DB: db.Get(),
	})
	paymentRepo := repository.NewPaymentRepository(repository.PaymentRepositoryConfig{
		DB: db.Get(),
	})
	reviewRepo := repository.NewReviewRepository(repository.ReviewRepositoryConfig{
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
	menuUsecase := usecase.NewMenuUsecase(usecase.MenuUsecaseConfig{
		MenuRepository: menuRepo,
	})
	promotionUsecase := usecase.NewPromotionUsecase(usecase.PromotionUsecaseConfig{
		PromotionRepository: promotionRepo,
	})
	cartUsecase := usecase.NewCartUsecase(usecase.CartUsecaseConfig{
		CartRepository:      cartRepo,
		UserRepository:      userRepo,
		MenuRepository:      menuRepo,
		PromotionRepository: promotionRepo,
	})
	orderUsecase := usecase.NewOrderUsecase(usecase.OrderUsecaseConfig{
		OrderRepository:  orderRepo,
		UserRepository:   userRepo,
		CartRepository:   cartRepo,
		CouponRepository: couponRepo,
	})
	paymentUsecase := usecase.NewPaymentUsecase(usecase.PaymentUsecaseConfig{
		PaymentRepository: paymentRepo,
	})
	reviewUsecase := usecase.NewReviewUsecase(usecase.ReviewUsecaseConfig{
		ReviewRepository: reviewRepo,
		UserRepository:   userRepo,
		OrderRepository:  orderRepo,
	})

	r := CreateRouter(RouterConfig{
		ExampleUsecase:   exampleUsecase,
		UserUsecase:      userUsecase,
		AuthUsecase:      authUsecase,
		MenuUsecase:      menuUsecase,
		PromotionUsecase: promotionUsecase,
		CartUsecase:      cartUsecase,
		OrderUsecase:     orderUsecase,
		PaymentUsecase:   paymentUsecase,
		ReviewUsecase:    reviewUsecase,
		AuthUtil:         authUtil,
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
