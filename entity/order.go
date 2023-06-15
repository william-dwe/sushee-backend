package entity

import (
	"time"

	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

const (
	ORDER_STATUS_PAYMENT   = "payment"
	ORDER_STATUS_PROCESSED = "processed"
	ORDER_STATUS_COMPLETED = "completed"
)

type Order struct {
	ID              uint           `json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
	UserId          int            `json:"user_id"`
	OrderDate       time.Time      `json:"order_date"`
	CouponId        *int           `json:"coupon_id"`
	PaymentOptionId int            `json:"payment_option_id"`
	OrderedMenus    []OrderedMenu  `json:"ordered_menus"`
	GrossAmount     float64        `json:"gross_amount"`
	DiscountAmount  float64        `json:"discount_amount"`
	NetAmount       float64        `json:"net_amount"`
	Status          string         `json:"status"`
}

type OrderedMenu struct {
	ID          uint           `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	OrderId     int            `json:"order_id"`
	MenuId      *int           `json:"menu_id"`
	Menu        Menu           `json:"menu"`
	PromotionId *int           `json:"promotion_id"`
	Quantity    int            `json:"quantity"`
	MenuOption  pgtype.JSONB   `json:"menu_option"`
	Review      Review         `json:"review"`
}

type DeliveryOrder struct {
	ID          uint           `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	OrderId     int            `json:"order_id"`
	DeliveredAt time.Time      `json:"delivered_at"`
}
