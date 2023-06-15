package repository

import (
	"sushee-backend/entity"
	"sushee-backend/httperror/domain"
	"sushee-backend/utils"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	AddReview(review *entity.Review) (*entity.Review, error)
	UpdateAvgReviewScoreByMenuId(MenuId int) error
}

type reviewRepositoryImpl struct {
	db *gorm.DB
}

type ReviewRepositoryConfig struct {
	DB *gorm.DB
}

func NewReviewRepository(c ReviewRepositoryConfig) ReviewRepository {
	return &reviewRepositoryImpl{
		db: c.DB,
	}
}

func (r *reviewRepositoryImpl) AddReview(review *entity.Review) (*entity.Review, error) {
	err := r.db.
		Create(review).
		Error
	if err != nil {
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"reviews_menu_id_fkey":         domain.ErrOrderRepoMenuNotFound,
				"reviews_ordered_menu_id_fkey": domain.ErrOrderRepoOrderedMenuNotFound,
			},
			domain.ErrOrderRepoInternal,
		)
		return nil, err
	}

	return review, nil
}

func (r *reviewRepositoryImpl) UpdateAvgReviewScoreByMenuId(MenuId int) error {
	sqAvgReview := r.db.
		Table("reviews").
		Select("AVG(rating)")

	err := r.db.
		Table("Menus").
		Where("id = (?)", MenuId).
		Update("avg_rating", sqAvgReview).
		Error
	if err != nil {
		return domain.ErrOrderRepoInternal
	}
	return nil
}
