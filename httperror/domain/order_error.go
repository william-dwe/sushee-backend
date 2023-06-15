package domain

import "sushee-backend/httperror"

var ErrOrderRepoInternal = httperror.InternalServerError("internal error")
var ErrOrderRepoPaymentOptionNotFound = httperror.BadRequestError("Payment option not found", "PAYMENT_OPTION_NOT_FOUND")
var ErrOrderRepoCouponNotFound = httperror.BadRequestError("Coupon not found", "COUPON_NOT_FOUND")
var ErrOrderRepoUserNotFound = httperror.BadRequestError("User not found", "USER_NOT_FOUND")
var ErrOrderRepoMenuNotFound = httperror.BadRequestError("Menu not found", "MENU_NOT_FOUND")
var ErrOrderRepoOrderNotFound = httperror.BadRequestError("Order not found", "ORDER_NOT_FOUND")
var ErrOrderRepoPromotionNotFound = httperror.BadRequestError("Promotion not found", "PROMOTION_NOT_FOUND")
var ErrOrderRepoOrderedMenuNotFound = httperror.BadRequestError("Ordered menu not found", "ORDERED_MENU_NOT_FOUND")

var ErrOrderUsecaseInternalError = httperror.InternalServerError("internal error")
var ErrOrderUsecaseOrderIsOrdered = httperror.BadRequestError("Order is ordered", "ORDER_IS_ORDERED")
var ErrOrderUsecaseUnauthorizedOrder = httperror.UnauthorizedError()
var ErrOrderUsecaseOrderEmpty = httperror.BadRequestError("Order is empty, please order something", "ORDER_EMPTY")
var ErrOrderUsecaseSomeCartIdNotFound = httperror.BadRequestError("Some cart id not found", "CART_ID_NOT_FOUND")

var ErrOrderHandlerInvalidRequestBody = httperror.BadRequestError("Invalid request body", "INVALID_REQUEST_BODY")
var ErrOrderHandlerIncorrectLimitRequest = httperror.BadRequestError("Incorrect limit request", "INCORRECT_LIMIT_REQUEST")
var ErrOrderHandlerIncorrectPageRequest = httperror.BadRequestError("Incorrect page request", "INCORRECT_PAGE_REQUEST")
