package repository

import (
	"sushee-backend/entity"
	"sushee-backend/httperror/domain"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	GetPaymentOption() (*[]entity.PaymentOption, error)
}

type paymentRepositoryImpl struct {
	db *gorm.DB
}

type PaymentRepositoryConfig struct {
	DB *gorm.DB
}

func NewPaymentRepository(c PaymentRepositoryConfig) PaymentRepository {
	return &paymentRepositoryImpl{
		db: c.DB,
	}
}

func (r *paymentRepositoryImpl) GetPaymentOption() (*[]entity.PaymentOption, error) {
	var payments []entity.PaymentOption
	err := r.db.
		Find(&payments).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrOrderRepoPaymentOptionNotFound
		}
		return nil, domain.ErrOrderRepoInternal
	}
	return &payments, err
}
