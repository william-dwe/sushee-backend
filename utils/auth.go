package utils

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"sushee-backend/config"
	"sushee-backend/entity"
	"sushee-backend/httperror"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
)

type AuthUtil interface {
	GenerateRefreshToken() (string, error)
	GenerateAccessToken(username string, scope string) (string, error)
	ValidateToken(encodedToken, signSecret string) (*jwt.Token, error)
	GenerateVerificationCode() (string, error)
	EmbedTokenOnContextCookie(c *gin.Context, refreshToken *string, accessToken *string, isUserLoggedIn *bool, appUrl string)
}

type authUtilImpl struct{}

func NewAuthUtil() AuthUtil {
	return &authUtilImpl{}
}

var c = config.Config.AuthConfig

type customAccessTokenClaims struct {
	User  entity.AuthTokenPayload `json:"user"`
	Scope string                  `json:"scope"`
	jwt.RegisteredClaims
}

func (a *authUtilImpl) GenerateAccessToken(username string, scope string) (string, error) {
	token := taylorAccessToken(username, scope)
	tokenStr, err := token.SignedString([]byte(c.HmacSecretAccessToken))

	return tokenStr, err
}

func taylorAccessToken(username, scope string) *jwt.Token {
	expirationLimit, _ := strconv.ParseInt(c.TimeLimitAccessToken, 10, 64)
	claims := &customAccessTokenClaims{
		entity.AuthTokenPayload{
			Username: username,
		},
		scope,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(expirationLimit))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.Config.AppConfig.Name,
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

type customRefreshTokenClaims struct {
	jwt.RegisteredClaims
}

func (a *authUtilImpl) GenerateRefreshToken() (string, error) {
	token := taylorRefreshToken()
	tokenString, err := token.SignedString([]byte(c.HmacSecretRefreshToken))

	return tokenString, err
}

func taylorRefreshToken() *jwt.Token {
	expirationLimit, _ := strconv.ParseInt(c.TimeLimitRefreshToken, 10, 64)
	claims := &customRefreshTokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(expirationLimit))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.Config.AppConfig.Name,
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func (a *authUtilImpl) ValidateToken(encodedToken, signSecret string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("invalid signature token")
		}
		return []byte(signSecret), nil
	})
}

func GetUserJWTContext(c *gin.Context) (*entity.AuthTokenPayload, error) {
	user, ok := c.Get("user")
	if !ok {
		log.Error().Msg("Not loggedin, cannot get user by token")
		return nil, httperror.UnauthorizedError()
	}

	userJson, _ := json.Marshal(user)

	var userPayload entity.AuthTokenPayload
	err := json.Unmarshal(userJson, &userPayload)
	if err != nil {
		return nil, httperror.UnauthorizedError()
	}

	userJWT := user.(entity.AuthTokenPayload)
	return &userJWT, nil
}

func GetScopeJWTContext(c *gin.Context) (string, error) {
	scope, ok := c.Get("scope")
	if !ok {
		return "", httperror.UnauthorizedError()
	}
	strScope, ok := scope.(string)
	if !ok {
		return "", httperror.UnauthorizedError()
	}
	return strScope, nil
}

const (
	verCodeChars  = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	verCodeDigits = 6
)

func (a *authUtilImpl) GenerateVerificationCode() (string, error) {
	buffer := make([]byte, verCodeDigits)
	rand.Seed(time.Now().UnixNano())
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(verCodeChars)
	for i := 0; i < verCodeDigits; i++ {
		buffer[i] = verCodeChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func (a *authUtilImpl) EmbedTokenOnContextCookie(c *gin.Context, refreshToken *string, accessToken *string, isUserLoggedIn *bool, appUrl string) {
	refreshTokenExpLimit, _ := strconv.Atoi(config.Config.AuthConfig.TimeLimitRefreshToken)
	accessTokenExpLimit, _ := strconv.Atoi(config.Config.AuthConfig.TimeLimitAccessToken)
	isRelease := config.Config.ENVConfig.Mode == config.ENV_MODE_RELEASE
	if isRelease {
		c.SetSameSite(http.SameSiteNoneMode)
	}
	if refreshToken != nil {
		c.SetCookie(
			"refresh_token",
			*refreshToken,
			refreshTokenExpLimit*60,
			"/",
			appUrl,
			isRelease,
			true,
		)
		if *refreshToken == "" {
			c.SetCookie(
				"refresh_token",
				"",
				-1,
				"/",
				appUrl,
				isRelease,
				true,
			)
		}
	}
	if accessToken != nil {
		c.SetCookie(
			"access_token",
			*accessToken,
			accessTokenExpLimit*60,
			"/",
			appUrl,
			isRelease,
			true,
		)
		if *accessToken == "" {
			c.SetCookie(
				"access_token",
				"",
				-1,
				"/",
				appUrl,
				isRelease,
				true,
			)
		}
	}

	if isUserLoggedIn != nil {
		c.SetCookie(
			"is_user_logged_in",
			strconv.FormatBool(*isUserLoggedIn),
			refreshTokenExpLimit*60,
			"/",
			appUrl,
			true,
			false,
		)
		if !*isUserLoggedIn {
			c.SetCookie(
				"is_user_logged_in",
				"",
				-1,
				"/",
				appUrl,
				true,
				false,
			)
		}
	}
}
