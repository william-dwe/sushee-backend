package dto

import (
	"sushee-backend/entity"

	"github.com/jackc/pgtype"
)

type CartResBody struct {
	ID             uint                      `json:"id"`
	UserId         int                       `json:"user_id"`
	MenuId         *int                      `json:"menu_id"`
	PromotionId    *int                      `json:"promotion_id"`
	Quantity       int                       `json:"quantity"`
	MenuOption     []entity.ChosenMenuOption `json:"menu_option"`
	IsOrdered      bool                      `json:"is_ordered"`
	PromotionPrice *float64                  `json:"promotion_price"`
}

type CartReqBody struct {
	MenuId      *int         `json:"menu_id,omitempty"`
	PromotionId *int         `json:"promotion_id,omitempty"`
	Quantity    int          `json:"quantity" binding:"required"`
	MenuOption  pgtype.JSONB `json:"menu_option,omitempty"`
}

type CartEditDetailsReqBody struct {
	Quantity   int          `json:"quantity,omitempty"`
	MenuOption pgtype.JSONB `json:"menu_option,omitempty"` // This might not working for now
}

type CartsResBody struct {
	Carts []CartResBody `json:"carts"`
}
