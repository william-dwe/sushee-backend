package entity

import (
	"time"

	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Cart struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	UserId         int
	MenuId         *int
	Menu           Menu
	PromotionId    *int
	Quantity       int
	MenuOption     pgtype.JSONB `gorm:"type:jsonb; default:'[]'"`
	IsOrdered      bool
	PromotionPrice *float64
}

type ChosenMenuOption struct {
	Title   string `json:"title" binding:"required"`
	Options string `json:"options" binding:"required"`
}
