package domain

import (
	"sushee-backend/httperror"
)

var ErrExampleAuth = httperror.UnauthorizedError()
var ErrExampleIdNotFound = httperror.BadRequestError("example ID not found", "DATA_NOT_FOUND")
var ErrCreateExample = httperror.InternalServerError("cannot create example record")
var ErrGetExample = httperror.InternalServerError("cannot get example record")
var ErrExampleUnexpected = httperror.InternalServerError("unexpected error occured in example process")
