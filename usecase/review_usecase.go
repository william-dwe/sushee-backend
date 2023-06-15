package usecase

import (
	"sushee-backend/dto"
	"sushee-backend/entity"
	"sushee-backend/httperror/domain"
	"sushee-backend/repository"
)

type ReviewUsecase interface {
	AddReview(username string, r *dto.ReviewAddReqBody) (*dto.ReviewResBody, error)
}

type reviewUsecaseImpl struct {
	orderRepository  repository.OrderRepository
	userRepository   repository.UserRepository
	reviewRepository repository.ReviewRepository
}

type ReviewUsecaseConfig struct {
	OrderRepository  repository.OrderRepository
	UserRepository   repository.UserRepository
	ReviewRepository repository.ReviewRepository
}

func NewReviewUsecase(c ReviewUsecaseConfig) ReviewUsecase {
	return &reviewUsecaseImpl{
		orderRepository:  c.OrderRepository,
		userRepository:   c.UserRepository,
		reviewRepository: c.ReviewRepository,
	}
}

func (u *reviewUsecaseImpl) AddReview(username string, r *dto.ReviewAddReqBody) (*dto.ReviewResBody, error) {
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, err
	}

	orderedMenu, err := u.orderRepository.GetOrderedMenuById(r.OrderedMenuId)
	if err != nil {
		return nil, err
	}

	order, err := u.orderRepository.GetOrderById(orderedMenu.OrderId)
	if err != nil {
		return nil, err
	}
	if order.UserId != int(user.ID) {
		return nil, domain.ErrReviewUsecaseUnauthorizedOrder
	}

	newReview := entity.Review{
		ReviewDescription: r.ReviewDescription,
		Rating:            r.Rating,
		OrderedMenuId:     r.OrderedMenuId,
		MenuId:            *orderedMenu.MenuId,
	}

	review, err := u.reviewRepository.AddReview(&newReview)
	if err != nil {
		return nil, err
	}

	err = u.reviewRepository.UpdateAvgReviewScoreByMenuId(newReview.MenuId)
	if err != nil {
		return nil, err
	}

	res := dto.ReviewResBody{
		ReviewDescription: review.ReviewDescription,
		Rating:            review.Rating,
		OrderedMenuId:     review.OrderedMenuId,
		MenuId:            review.MenuId,
	}

	return &res, nil
}
