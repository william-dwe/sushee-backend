package usecase

import (
	"sushee-backend/dto"
	"sushee-backend/entity"
	"sushee-backend/repository"
)

type ExampleUsecase interface {
	ExampleProcess(dto.ExampleReqDTO) (*dto.ExampleResDTO, error)
}

type ExampleUsecaseConfig struct {
	ExampleRepository repository.ExampleRepository
}

type exampleUsecaseImpl struct {
	exampleRepository repository.ExampleRepository
}

func NewExampleUsecase(c ExampleUsecaseConfig) ExampleUsecase {
	return &exampleUsecaseImpl{
		exampleRepository: c.ExampleRepository,
	}
}

func (u *exampleUsecaseImpl) ExampleProcess(input dto.ExampleReqDTO) (*dto.ExampleResDTO, error) {
	//do something
	createExample := entity.Example{
		//set value
		Qty:  10,
		Name: input.ExampleField,
	}

	err := u.exampleRepository.Store(createExample)
	if err != nil {
		return nil, err
	}

	return &dto.ExampleResDTO{}, nil
}
