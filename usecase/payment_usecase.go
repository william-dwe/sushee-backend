package usecase

import (
	"sushee-backend/dto"
	"sushee-backend/repository"
)

type PaymentUsecase interface {
	GetPaymentOption() (*dto.PaymentOptionsResBody, error)
}

type paymentUsecaseImpl struct {
	paymentRepository repository.PaymentRepository
}

type PaymentUsecaseConfig struct {
	PaymentRepository repository.PaymentRepository
}

func NewPaymentUsecase(c PaymentUsecaseConfig) PaymentUsecase {
	return &paymentUsecaseImpl{
		paymentRepository: c.PaymentRepository,
	}
}

func (u *paymentUsecaseImpl) GetPaymentOption() (*dto.PaymentOptionsResBody, error) {
	paymentOptions, err := u.paymentRepository.GetPaymentOption()
	if err != nil {
		return nil, err
	}

	poSlice := []dto.PaymentOptionResBody{}
	for _, po := range *paymentOptions {
		poSlice = append(poSlice, dto.PaymentOptionResBody{
			ID:          int(po.ID),
			PaymentName: po.PaymentName,
		})
	}

	res := dto.PaymentOptionsResBody{
		PaymentOptions: poSlice,
	}

	return &res, nil
}
