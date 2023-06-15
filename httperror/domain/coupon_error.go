package domain

import "sushee-backend/httperror"

var ErrCouponRepoInternal = httperror.InternalServerError("internal error")
var ErrCoponRepoAdminNotFound = httperror.BadRequestError("Admin not found", "ADMIN_NOT_FOUND")
var ErrCouponRepoCouponNotFound = httperror.BadRequestError("Coupon not found", "COUPON_NOT_FOUND")
var ErrCouponRepoUserNotFound = httperror.BadRequestError("User not found", "USER_NOT_FOUND")
