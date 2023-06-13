package domain

import "sushee-backend/httperror"

var ErrUploadFile = httperror.BadRequestError("upload file error", "UPLOAD_FILE_ERROR")
var ErrDeleteFile = httperror.BadRequestError("delete file error", "DELETE_FILE_ERROR")

var ErrFileSizeExceedLimit = httperror.BadRequestError("file size exceed limit", "FILE_SIZE_EXCEED_LIMIT")
var ErrFileTypeNotAllowed = httperror.BadRequestError("file type not allowed", "FILE_TYPE_NOT_ALLOWED")
