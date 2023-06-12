package repository

import (
	"sushee-backend/entity"
	"sushee-backend/httperror/domain"

	"gorm.io/gorm"
)

type ExampleRepository interface {
	Store(entity.Example) error
	GetByID(uint) (*entity.Example, error)
}

type ExampleRepositoryConfig struct {
	DB *gorm.DB
}

type exampleRepositoryImpl struct {
	db *gorm.DB
}

func NewExampleRepository(c ExampleRepositoryConfig) ExampleRepository {
	return &exampleRepositoryImpl{
		db: c.DB,
	}
}

func (r *exampleRepositoryImpl) Store(input entity.Example) error {
	err := r.db.Create(&input).Error
	if err != nil {
		return domain.ErrCreateExample
	}

	return nil
}

func (r *exampleRepositoryImpl) GetByID(id uint) (*entity.Example, error) {
	var example entity.Example
	err := r.db.First(&example, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrExampleIdNotFound
		}
		//check another error or if is not masked, return internal error
		return nil, domain.ErrGetExample
	}

	return &example, nil
}
