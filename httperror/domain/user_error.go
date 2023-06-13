package domain

import "sushee-backend/httperror"

var ErrUserRepoInternal = httperror.InternalServerError("Internal error")
var ErrUserRepoUserNotFound = httperror.BadRequestError("User not found", "USER_NOT_FOUND")
var ErrUserRepoRoleNotFound = httperror.BadRequestError("Role not found", "ROLE_NOT_FOUND")
var ErrUserRepoEmailAlreadyExist = httperror.BadRequestError("Email already exist", "EMAIL_ALREADY_EXIST")
var ErrUserRepoPhoneAlreadyExist = httperror.BadRequestError("Phone already exist", "PHONE_ALREADY_EXIST")
var ErrUserRepoInvalidPhoneFormat = httperror.BadRequestError("Invalid phone format", "INVALID_PHONE_FORMAT")
var ErrUserRepoUsernameAlreadyExist = httperror.BadRequestError("Username already exist", "USERNAME_ALREADY_EXIST")
var ErrUserRepoInvalidUsernameFormat = httperror.BadRequestError("Invalid username format", "INVALID_USERNAME_FORMAT")
var ErrUserRepoDetailRoleNotFound = httperror.BadRequestError("Detail role not found", "DETAIL_ROLE_NOT_FOUND")

var ErrUserHandlerUpdateIncompleteReqBody = httperror.BadRequestError("invalid input", "UPDATE_INPUT_INCOMPLETE")
