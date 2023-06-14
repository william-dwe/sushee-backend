package repository

import (
	"sushee-backend/entity"
	"sushee-backend/httperror/domain"
	"sushee-backend/utils"
	"time"

	"gorm.io/gorm"
)

type PromotionRepository interface {
	GetPromotionMenu() (*[]entity.Promotion, error)
	GetAndValidatePromoMenu(menuId, promoId int) (*entity.PromoMenu, error)
}

type promotionRepositoryImpl struct {
	db *gorm.DB
}

type PromotionRepositoryConfig struct {
	DB *gorm.DB
}

func NewPromotionRepository(c PromotionRepositoryConfig) PromotionRepository {
	return &promotionRepositoryImpl{
		db: c.DB,
	}
}

func (r *promotionRepositoryImpl) GetPromotionMenu() (*[]entity.Promotion, error) {
	var promotions []entity.Promotion
	err := r.db.
		Model(&entity.Promotion{}).
		Preload("PromoMenus.Menu").
		Where("? between started_at and expired_at", time.Now()).
		Find(&promotions).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrMenuRepoNoPromotionFound
		}

		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"promo_menus_menu_id_fkey":      domain.ErrMenuRepoMenuNotFound,
				"promo_menus_promotion_id_fkey": domain.ErrMenuRepoPromotionNotFound,
			},
			domain.ErrMenuRepoInternal,
		)
		return nil, err
	}
	return &promotions, nil
}

func (r *promotionRepositoryImpl) GetAndValidatePromoMenu(menuId, promoId int) (*entity.PromoMenu, error) {
	var c entity.PromoMenu
	err := r.db.
		Where("menu_id = ? AND promotion_id = ?", menuId, promoId).
		Find(&c).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrPromotionMenuRepoInvalidPromo
		}

		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"promo_menus_menu_id_fkey":      domain.ErrMenuRepoMenuNotFound,
				"promo_menus_promotion_id_fkey": domain.ErrMenuRepoPromotionNotFound,
			},
			domain.ErrMenuRepoInternal,
		)
		return nil, err
	}
	return &c, nil
}
