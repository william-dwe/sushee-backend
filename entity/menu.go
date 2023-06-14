package entity

import (
	"time"

	"github.com/jackc/pgtype"
	"gorm.io/gorm"
)

type Menu struct {
	ID                uint
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
	MenuName          string
	AvgRating         float64
	NumberOfFavorites int
	Price             float64
	MenuPhoto         string
	CategoryId        int
	Customization     pgtype.JSONB `gorm:"type:jsonb;;default:'[]'"`
	Reviews           []Review
}

type MenuCustomization struct {
	Title   string   `json:"title"`
	Type    string   `json:"type"`
	Options []string `json:"options"`
}

type MenuCategory struct {
	gorm.Model
	CategoryName string
}

type PromoMenu struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	PromotionId    int
	MenuId         int
	Menu           Menu `gorm:"foreignKey:MenuId;references:ID"`
	PromotionPrice float64
}

type Promotion struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	AdminId        int
	Name           string
	Description    string
	PromotionPhoto string
	DiscountRate   float64
	StartAt        time.Time
	ExpiredAt      time.Time
	PromoMenus     []PromoMenu `gorm:"foreignKey:PromotionId;references:ID"`
}
