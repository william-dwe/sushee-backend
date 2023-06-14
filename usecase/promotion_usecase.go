package usecase

import (
	"sushee-backend/entity"
	"sushee-backend/repository"
)

type PromotionUsecase interface {
	GetPromotion() (*[]entity.Promotion, error)
}

type promotionUsecaseImpl struct {
	promotionRepository repository.PromotionRepository
}

type PromotionUsecaseConfig struct {
	PromotionRepository repository.PromotionRepository
}

func NewPromotionUsecase(c PromotionUsecaseConfig) PromotionUsecase {
	return &promotionUsecaseImpl{
		promotionRepository: c.PromotionRepository,
	}
}

func (u *promotionUsecaseImpl) GetPromotion() (*[]entity.Promotion, error) {
	menus, err := u.promotionRepository.GetPromotionMenu()
	if err != nil {
		return nil, err
	}

	return menus, nil
}
