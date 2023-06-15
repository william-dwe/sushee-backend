package entity

import (
	"time"

	"gorm.io/gorm"
)

type PaymentOption struct {
	ID          uint           `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	PaymentName string         `json:"payment_name"`
}

func (PaymentOption) TableName() string {
	return "payment_options"
}
