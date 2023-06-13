package utils

import (
	"fmt"

	"sushee-backend/httperror"
	"sushee-backend/httperror/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ShouldBindJsonWithValidation(c *gin.Context, dto any) error {
	if err := c.ShouldBindJSON(dto); err != nil {
		if validatorErrHandler(err) != nil {
			return err
		}
	}

	return nil
}

func ShouldBindWithValidation(c *gin.Context, dto any) error {
	if err := c.ShouldBind(dto); err != nil {
		if validatorErrHandler(domain.ErrReqHandlerValidatorInvalid) != nil {
			return err
		}
	}

	return nil
}

func validatorErrHandler(err error) error {
	errTag, cvtErr := err.(validator.ValidationErrors)
	if !cvtErr {
		return err
	}
	if errTag != nil {
		return httperror.BadRequestError(fmt.Sprintf("check the requested param: %s must be %s %s",
			ToSnakeCase(errTag[0].StructField()),
			errTag[0].ActualTag(),
			errTag[0].Param()), "DATA_NOT_VALID")
	}
	return nil
}
