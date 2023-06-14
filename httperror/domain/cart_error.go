package domain

import "sushee-backend/httperror"

var ErrCartRepoInternal = httperror.InternalServerError("internal error")
var ErrCartRepoAddItemToCartInvalidMenuId = httperror.BadRequestError("invalid menu id", "INVALID_MENU_ID")
var ErrCartRepoAddItemToCartInvalidPromotionId = httperror.BadRequestError("invalid promotion id", "INVALID_PROMOTION_ID")
var ErrCartRepoAddItemToCartInvalidUserId = httperror.BadRequestError("invalid user id", "INVALID_USER_ID")
var ErrCartRepoNoCartFound = httperror.BadRequestError("no cart found", "NO_CART_FOUND")
var ErrCartRepoCartIdNotFound = httperror.BadRequestError("cart id not found", "CART_ID_NOT_FOUND")

var ErrCartUsecaseMarshallMenuOption = httperror.BadRequestError("menu option format incorrect", "MARSHALL_MENU_OPTION_ERROR")
var ErrCartUsecaseUnmarshallMenuOption = httperror.InternalServerError("unmarshall menu option error")
var ErrCartUsecaseInvalidCartOwner = httperror.BadRequestError("invalid cart owner", "INVALID_CART_OWNER")
var ErrCartUsecaseMenuOrPromotionMenuNotFound = httperror.BadRequestError("menu or promotion menu not found", "MENU_OR_PROMOTION_MENU_NOT_FOUND")

var ErrCartHandlerIncompleteReqBody = httperror.BadRequestError("incomplete request body", "INCOMPLETE_REQUEST_BODY")
var ErrCartHandlerInvalidCartId = httperror.BadRequestError("invalid cart id", "INVALID_CART_ID")
