package entity

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	ID             uint           `json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	AdminId        int            `json:"admin_id"`
	Description    string         `json:"description"`
	DiscountAmount float64        `json:"discount_amount"`
	QuotaInitial   int            `json:"quota_initial"`
	QuotaLeft      int            `json:"quota_left"`
}

type UserCoupon struct {
	ID             uint           `json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	UserId         int            `json:"user_id"`
	CouponId       int            `json:"coupon_id"`
	CouponCode     string         `json:"coupon_code"`
	DiscountAmount float64        `json:"discount_amount"`
}
