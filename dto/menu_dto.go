package dto

import (
	"sushee-backend/entity"

	"github.com/jackc/pgtype"
)

type MenuAddReqBody struct {
	MenuName      string       `json:"menu_name"`
	Price         float64      `json:"price"`
	MenuPhoto     string       `json:"menu_photo"`
	CategoryId    int          `json:"category_id"`
	Customization pgtype.JSONB `json:"customization"`
}

type MenuQuery struct {
	Search           string `json:"search"`
	SortBy           string `json:"sort_by"`
	FilterByCategory string `json:"filter_by_category"`
	Sort             string `json:"sort"`
	Limit            int    `json:"limit"`
	Page             int    `json:"page"`
}

type MenuItem struct {
	MenuName          string                     `json:"menu_name"`
	AvgRating         float64                    `json:"avg_rating"`
	NumberOfFavorites int                        `json:"number_of_favorites"`
	Price             float64                    `json:"price"`
	MenuPhoto         string                     `json:"menu_photo"`
	CategoryId        int                        `json:"category_id"`
	Customization     []entity.MenuCustomization `json:"customization"`
}

type MenusResBody struct {
	Menus       []MenuItem `json:"menus"`
	CurrentPage int        `json:"current_page"`
	MaxPage     int        `json:"max_page"`
}

type PromotionResBody struct {
	Promotions []entity.Promotion `json:"promotions"`
}
