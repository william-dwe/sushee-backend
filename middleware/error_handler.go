package middleware

import (
	"sushee-backend/httperror"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	firstErr := c.Errors[0].Err
	appErr, isAppError := firstErr.(httperror.AppError)
	if isAppError {
		c.AbortWithStatusJSON(appErr.StatusCode, appErr)
		return
	}
	serverErr := httperror.InternalServerError(firstErr.Error())
	c.AbortWithStatusJSON(serverErr.StatusCode, serverErr)
}
