package repository

import (
	"sushee-backend/entity"
	"sushee-backend/httperror/domain"
	"sushee-backend/utils"

	"gorm.io/gorm"
)

type CouponRepository interface {
	AddCoupon(coupon *entity.Coupon) (*entity.Coupon, error)
	GetCoupon() (*[]entity.Coupon, error)
	GetCouponById(couponId int) (*entity.Coupon, error)
	UpdateCouponById(couponId int, newCoupon *entity.Coupon) error
	DeleteCouponById(couponId int) (*entity.Coupon, error)
	AddUserCoupon(userCoupon *entity.UserCoupon) (*entity.UserCoupon, error)
	GetUserCouponByUsername(username string) (*[]entity.UserCoupon, int, error)
	GetUserCouponByCouponCode(userId int, couponCode string) (*entity.UserCoupon, error)
}

type CouponRepositoryImpl struct {
	db *gorm.DB
}

type CouponRepositoryConfig struct {
	DB *gorm.DB
}

func NewCouponRepository(c CouponRepositoryConfig) CouponRepository {
	return &CouponRepositoryImpl{
		db: c.DB,
	}
}

func (r *CouponRepositoryImpl) AddCoupon(coupon *entity.Coupon) (*entity.Coupon, error) {
	err := r.db.
		Create(&coupon).Error
	if err != nil {
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"coupons_admin_id_fkey": domain.ErrCoponRepoAdminNotFound,
			},
			domain.ErrCouponRepoInternal,
		)
	}
	return coupon, err
}

func (r *CouponRepositoryImpl) GetCoupon() (*[]entity.Coupon, error) {
	var coupon []entity.Coupon
	err := r.db.
		Find(&coupon).Error
	if err != nil {
		return nil, domain.ErrCouponRepoInternal
	}
	return &coupon, nil
}

func (r *CouponRepositoryImpl) GetCouponById(couponId int) (*entity.Coupon, error) {
	var coupon entity.Coupon

	err := r.db.
		Where("id = (?)", couponId).Debug().
		First(&coupon).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrCouponRepoCouponNotFound
		}
		return nil, domain.ErrCouponRepoInternal
	}

	return &coupon, nil
}

func (r *CouponRepositoryImpl) UpdateCouponById(couponId int, newCoupon *entity.Coupon) error {
	err := r.db.
		Where("id = (?)", couponId).
		Updates(newCoupon).
		Debug().Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.ErrCouponRepoCouponNotFound
		}

		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"coupons_admin_id_fkey": domain.ErrCoponRepoAdminNotFound,
			},
			domain.ErrCouponRepoInternal,
		)
	}
	return err
}

func (r *CouponRepositoryImpl) DeleteCouponById(couponId int) (*entity.Coupon, error) {
	var coupon entity.Coupon
	err := r.db.
		Where("id = (?)", couponId).
		Delete(&coupon).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrCouponRepoCouponNotFound
		}
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"coupons_admin_id_fkey": domain.ErrCoponRepoAdminNotFound,
			},
			domain.ErrCouponRepoInternal,
		)
	}
	return &coupon, err
}

func (r *CouponRepositoryImpl) AddUserCoupon(userCoupon *entity.UserCoupon) (*entity.UserCoupon, error) {
	err := r.db.
		Create(&userCoupon).
		Error
	if err != nil {
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"user_coupons_coupon_id_fkey": domain.ErrCouponRepoCouponNotFound,
				"user_coupons_user_id_fkey":   domain.ErrCouponRepoUserNotFound,
			},
			domain.ErrCouponRepoInternal,
		)
	}
	return userCoupon, err
}

func (r *CouponRepositoryImpl) GetUserCouponByUsername(username string) (*[]entity.UserCoupon, int, error) {
	var coupon []entity.UserCoupon
	userSQ := r.db.
		Table("users").
		Select("id").
		Where("username = (?)", username)
	q := r.db.
		Where("user_id in (?)", userSQ).
		Find(&coupon)
	err := q.Error
	if err != nil {
		return nil, 0, domain.ErrCouponRepoInternal
	}
	return &coupon, int(q.RowsAffected), q.Error
}

func (r *CouponRepositoryImpl) GetUserCouponByCouponCode(userId int, couponCode string) (*entity.UserCoupon, error) {
	var coupon entity.UserCoupon
	q := r.db.
		Where("user_id in (?) AND coupon_code = (?)", userId, couponCode).
		First(&coupon)
	err := q.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrCouponRepoCouponNotFound
		}
		return nil, domain.ErrCouponRepoInternal
	}
	return &coupon, q.Error
}
