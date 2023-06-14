package handler

import (
	"strconv"
	"sushee-backend/dto"
	"sushee-backend/entity"
	"sushee-backend/httperror/domain"
	"sushee-backend/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ShowMenu(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil {
		_ = c.Error(domain.ErrMenuHandlerInvalidLimitRequest)
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		_ = c.Error(domain.ErrMenuHandlerInvalidPageRequest)
		return
	}

	q := dto.MenuQuery{
		Search:           c.DefaultQuery("s", "%"),
		SortBy:           c.DefaultQuery("sortBy", "category_id"),
		FilterByCategory: c.DefaultQuery("filterByCategory", ""),
		Sort:             c.DefaultQuery("sort", "desc"),
		Limit:            limit,
		Page:             page,
	}

	m, err := h.menuUsecase.GetMenu(q)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_SHOW_MENU",
		Message: "success show menu",
		Data:    m,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) ShowPromotion(c *gin.Context) {
	t, err := h.menuUsecase.GetPromotion()
	if err != nil {
		_ = c.Error(err)
		return
	}

	respBody := dto.PromotionResBody{
		Promotions: *t,
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_SHOW_PROMOTION",
		Message: "success show promotion",
		Data:    respBody,
	}

	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) AddMenu(c *gin.Context) {
	var reqBody dto.MenuAddReqBody
	if err := utils.ShouldBindJsonWithValidation(c, &reqBody); err != nil {
		_ = c.Error(err)
		return
	}

	menu, err := h.menuUsecase.AddMenu(&reqBody)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_ADD_MENU",
		Message: "success add menu",
		Data:    menu,
	}
	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) UpdateMenu(c *gin.Context) {
	menuId, err := strconv.Atoi(c.Param("menuId"))
	if err != nil {
		_ = c.Error(domain.ErrMenuHandlerInvalidMenuId)
		return
	}
	var reqBody dto.MenuAddReqBody
	if err := c.BindJSON(&reqBody); err != nil {
		_ = c.Error(domain.ErrMenuHandlerInvalidMenuName)
		return
	}

	newMenu := entity.Menu{
		MenuName:      reqBody.MenuName,
		Price:         reqBody.Price,
		MenuPhoto:     reqBody.MenuPhoto,
		CategoryId:    reqBody.CategoryId,
		Customization: reqBody.Customization,
	}

	menu, err := h.menuUsecase.UpdateMenuByMenuId(menuId, &newMenu)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_UPDATE_MENU",
		Message: "success update menu",
		Data:    menu,
	}
	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) DeleteMenu(c *gin.Context) {
	menuId, err := strconv.Atoi(c.Param("menuId"))
	if err != nil {
		_ = c.Error(domain.ErrMenuHandlerInvalidMenuId)
		return
	}

	err = h.menuUsecase.DeleteMenuByMenuId(menuId)
	if err != nil {
		_ = c.Error(err)
		return
	}
	res := dto.ResponseStruct{
		Code:    "SUCCESS_DELETE_MENU",
		Message: "success delete menu",
		Data:    nil,
	}
	utils.ResponseSuccessJSON(c, res)
}

func (h *Handler) GetMenuDetail(c *gin.Context) {
	menuId, err := strconv.Atoi(c.Param("menuId"))
	if err != nil {
		_ = c.Error(domain.ErrMenuHandlerInvalidMenuId)
		return
	}

	menu, err := h.menuUsecase.GetMenuDetailByMenuId(menuId)
	if err != nil {
		_ = c.Error(err)
		return
	}

	res := dto.ResponseStruct{
		Code:    "SUCCESS_GET_MENU_DETAIL",
		Message: "success get menu detail",
		Data:    menu,
	}

	utils.ResponseSuccessJSON(c, res)
}
