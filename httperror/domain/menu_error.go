package domain

import "sushee-backend/httperror"

var ErrMenuRepoInternal = httperror.InternalServerError("Internal error")
var ErrMenuRepoCategoryNotFound = httperror.BadRequestError("Category not found", "CATEGORY_NOT_FOUND")
var ErrMenuRepoPromotionNotFound = httperror.BadRequestError("Promotion not found", "PROMOTION_NOT_FOUND")
var ErrMenuRepoMenuNotFound = httperror.BadRequestError("Menu not found", "MENU_NOT_FOUND")
var ErrMenuRepoNoMenuFound = httperror.BadRequestError("No menu found", "NO_MENU_FOUND")
var ErrMenuRepoNoPromotionFound = httperror.BadRequestError("No promotion found", "NO_PROMOTION_FOUND")
var ErrPromotionMenuRepoInvalidPromo = httperror.BadRequestError("Invalid promotion", "INVALID_PROMOTION")

var ErrMenuUsecaseMarshallCustomizese = httperror.BadRequestError("Input customization in the correct format", "MARSHALL_CUSTOMIZESE_ERROR")
var ErrMenuUsecaseUnmarshallCustomizese = httperror.BadRequestError("Unmarshall customizese error", "UNMARSHALL_CUSTOMIZESE_ERROR")

var ErrMenuHandlerInvalidLimitRequest = httperror.BadRequestError("Invalid limit request", "INVALID_LIMIT_REQUEST")
var ErrMenuHandlerInvalidPageRequest = httperror.BadRequestError("Invalid page request", "INVALID_PAGE_REQUEST")
var ErrMenuHandlerInvalidMenuId = httperror.BadRequestError("Invalid menu id request", "INVALID_MENU_ID_REQUEST")
var ErrMenuHandlerInvalidMenuName = httperror.BadRequestError("Invalid menu name request", "INVALID_MENU_NAME_REQUEST")
