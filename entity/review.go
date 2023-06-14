package entity

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID                uint           `json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at"`
	ReviewDescription string         `json:"review_description"`
	Rating            float64        `json:"rating"`
	OrderedMenuId     int            `json:"ordered_menu_id"`
	MenuId            int            `json:"menu_id"`
}

type ReviewAddReqBody struct {
	ReviewDescription string  `json:"review_description"`
	Rating            float64 `json:"rating"`
	OrderedMenuId     int     `json:"ordered_menu_id"`
}
