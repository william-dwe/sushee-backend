package domain

import "sushee-backend/httperror"

var ErrAuthHandlerRegisterIncompleteReqBody = httperror.BadRequestError("should provide complete register data", "REGISTER_INPUT_INCOMPLETE")
var ErrAuthHandlerLoginIncompleteReqBody = httperror.BadRequestError("should provide identifier and password", "LOGIN_INPUT_INCOMPLETE")
var ErrAuthHandlerLogoutTokenNotFound = httperror.UnauthorizedError()
var ErrAuthHandlerRefreshTokenNotFound = httperror.UnauthorizedError()

var ErrAuthRepoInternal = httperror.InternalServerError("internal error")
var ErrAuthRepoAddAuthSessionUserNotFound = httperror.BadRequestError("user not found", "USER_NOT_FOUND")
var ErrAuthRepoAuthSessionNotFound = httperror.BadRequestError("auth session not found", "AUTH_SESSION_NOT_FOUND")
