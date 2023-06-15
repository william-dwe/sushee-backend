package dto

import (
	"sushee-backend/entity"
	"time"
)

type OrderHistoryResBody struct {
	Orders []entity.Order `json:"orders"`
}

type OrderStatusUpdateReqBody struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type OrderReqBody struct {
	CartIdList      []int  `json:"cart_id_list"`
	PaymentOptionId int    `json:"payment_option_id"`
	CouponCode      string `json:"coupon_code,omitempty"`
}

type OrdersResBody struct {
	Orders      []OrderResBody `json:"orders"`
	CurrentPage int            `json:"current_page"`
	MaxPage     int            `json:"max_page"`
}

type OrderResBody struct {
	ID              int                  `json:"id"`
	UserId          int                  `json:"user_id"`
	OrderDate       time.Time            `json:"order_date"`
	CouponId        int                  `json:"coupon_id"`
	PaymentOptionId int                  `json:"payment_option_id"`
	OrderedMenus    []OrderedMenuResBody `json:"ordered_menus"`
	GrossAmount     float64              `json:"gross_amount"`
	DiscountAmount  float64              `json:"discount_amount"`
	NetAmount       float64              `json:"net_amount"`
	Status          string               `json:"status"`
}

type OrderedMenuResBody struct {
	MenuName      string                    `json:"menu_name"`
	MenuId        int                       `json:"menu_id"`
	PromotionId   int                       `json:"promotion_id"`
	Quantity      int                       `json:"quantity"`
	Customization []entity.ChosenMenuOption `json:"customization"`
}

type OrderHistoryQuery struct {
	Search           string
	SortBy           string
	FilterByCategory string
	Sort             string
	Limit            int
	Page             int
}

type OrderStatusQuery struct {
	Search         string
	SortBy         string
	FilterByStatus string
	Sort           string
	Limit          int
	Page           int
}
