package dto

type ReviewAddReqBody struct {
	ReviewDescription string  `json:"review_description"`
	Rating            float64 `json:"rating"`
	OrderedMenuId     int     `json:"ordered_menu_id"`
}

type ReviewResBody struct {
	ReviewDescription string  `json:"review_description"`
	Rating            float64 `json:"rating"`
	OrderedMenuId     int     `json:"ordered_menu_id"`
	MenuId            int     `json:"menu_id"`
}
