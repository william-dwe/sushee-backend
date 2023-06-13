package middleware

import (
	"encoding/json"
	"sushee-backend/config"
	"sushee-backend/entity"
	"sushee-backend/httperror"
	"sushee-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authenticate(c *gin.Context) {
	conf := config.Config.AuthConfig
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		utils.AbortWithError(c, httperror.UnauthorizedError())
		return
	}

	a := utils.NewAuthUtil()
	token, err := a.ValidateToken(accessToken, conf.HmacSecretAccessToken)
	if err != nil || !token.Valid {
		utils.AbortWithError(c, httperror.UnauthorizedError())
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		utils.AbortWithError(c, httperror.UnauthorizedError())
		return
	}

	userJson, _ := json.Marshal(claims["user"])
	var userPayload entity.AuthTokenPayload
	err = json.Unmarshal(userJson, &userPayload)
	if err != nil {
		utils.AbortWithError(c, httperror.UnauthorizedError())
		return
	}

	c.Set("user", userPayload)
	c.Set("scope", claims["scope"])
}
