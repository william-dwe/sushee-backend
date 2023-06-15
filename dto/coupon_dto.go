package dto

import "sushee-backend/entity"

type CouponAddReqBody struct {
	Description    string  `json:"description"`
	DiscountAmount float64 `json:"discount_amount"`
	QuotaInitial   int     `json:"quota_initial"`
	QuotaLeft      int     `json:"quota_left"`
}

type CouponEditReqBody struct {
	Description    string  `json:"description,omitempty"`
	DiscountAmount float64 `json:"discount_amount,omitempty"`
	QuotaInitial   int     `json:"quota_initial,omitempty"`
	QuotaLeft      int     `json:"quota_left,omitempty"`
}

type CouponResBody struct {
	Coupons []entity.Coupon `json:"coupons"`
}

type UserCouponAddReqBody struct {
	UserId   int `json:"user_id"`
	CouponId int `json:"coupon_id"`
}

type UserCouponResBody struct {
	UserCoupons []entity.UserCoupon `json:"user_coupons"`
}
