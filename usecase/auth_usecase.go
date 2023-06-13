package usecase

import (
	"strconv"
	"strings"
	"sushee-backend/httperror"
	"sushee-backend/utils"
	"time"

	"sushee-backend/config"
	"sushee-backend/dto"
	"sushee-backend/entity"
	"sushee-backend/repository"
)

type AuthUsecase interface {
	Register(*dto.UserRegisterReqBody) (*dto.UserRegisterResBody, error)
	Login(*dto.UserLoginReqBody) (res *dto.UserLoginResBody, accessToken string, refreshToken string, err error)
	Logout(refreshToken string) (*dto.UserLogoutResBody, error)
	Refresh(refreshToken string) (res *dto.UserLoginResBody, accessToken string, err error)
}

type authUsecaseImpl struct {
	authRepository repository.AuthRepository
	userRepository repository.UserRepository
	authUtil       utils.AuthUtil
}

type AuthUsecaseConfig struct {
	AuthRepository repository.AuthRepository
	UserRepository repository.UserRepository
	AuthUtil       utils.AuthUtil
}

func NewAuthUsecase(c AuthUsecaseConfig) AuthUsecase {
	return &authUsecaseImpl{
		authRepository: c.AuthRepository,
		userRepository: c.UserRepository,
		authUtil:       c.AuthUtil,
	}
}

func (u *authUsecaseImpl) Register(reqBody *dto.UserRegisterReqBody) (*dto.UserRegisterResBody, error) {
	isDuplicate, err := u.userRepository.CheckDuplicatePhone(reqBody.Phone)
	if err != nil || isDuplicate {
		return nil, err
	}

	initialProfilePicture := ""
	initialPlayAttempt := 0
	defaultRoleId := 1
	hashedPass, _ := utils.HashAndSalt(reqBody.Password)
	validReqNewUser := entity.User{
		FullName:       reqBody.FullName,
		Phone:          reqBody.Phone,
		Email:          reqBody.Email,
		Username:       reqBody.Username,
		Password:       hashedPass,
		RegisterDate:   time.Now(),
		ProfilePicture: initialProfilePicture,
		PlayAttempt:    initialPlayAttempt,
		RoleId:         defaultRoleId,
	}

	newUser, err := u.userRepository.AddNewUser(&validReqNewUser)
	if err != nil {
		return nil, err
	}

	validResNewUser := dto.UserRegisterResBody{
		FullName:     newUser.FullName,
		Phone:        newUser.Phone,
		Email:        newUser.Email,
		Username:     newUser.Username,
		RegisterDate: newUser.RegisterDate,
	}

	return &validResNewUser, nil
}

func (u *authUsecaseImpl) Login(req *dto.UserLoginReqBody) (*dto.UserLoginResBody, string, string, error) {
	var user *entity.User
	var err error
	user, err = u.userRepository.GetUserByEmailOrUsername(strings.ToLower(req.Identifier))
	if err != nil {
		return nil, "", "", err
	}
	role, err := u.userRepository.GetDetailRole(user.RoleId)
	if err != nil {
		return nil, "", "", err
	}

	if !utils.ValidateHash(user.Password, req.Password) {
		return nil, "", "", httperror.UnauthorizedError()
	}

	accessTokenStr, err := u.authUtil.GenerateAccessToken(user.Username, role.RoleName)
	if err != nil {
		return nil, "", "", err
	}
	refreshTokenStr, err := u.authUtil.GenerateRefreshToken()
	if err != nil {
		return nil, "", "", err
	}

	expirationLimit, _ := strconv.ParseInt(config.Config.AuthConfig.TimeLimitRefreshToken, 10, 64)
	session := entity.AuthSession{
		RefreshToken: refreshTokenStr,
		UserId:       user.ID,
		ExpiredAt:    time.Now().Add(time.Second * time.Duration(expirationLimit)),
	}
	_, err = u.authRepository.AddAuthSession(&session)
	if err != nil {
		return nil, "", "", err
	}

	token := dto.UserLoginResBody{
		AccessToken: accessTokenStr,
		Username:    user.Username,
		Email:       user.Email,
	}
	return &token, accessTokenStr, refreshTokenStr, err
}

func (u *authUsecaseImpl) Refresh(refreshToken string) (*dto.UserLoginResBody, string, error) {
	var err error
	a := utils.NewAuthUtil()
	_, err = a.ValidateToken(refreshToken, config.Config.AuthConfig.HmacSecretRefreshToken)
	if err != nil {
		return nil, "", err
	}

	session, err := u.authRepository.GetAuthSessionByRefreshToken(refreshToken)
	if err != nil {
		return nil, "", err
	}
	if time.Now().After(session.ExpiredAt) {
		_ = u.authRepository.DeleteAuthSessionById(session.ID)
		return nil, "", httperror.UnauthorizedError()
	}

	user, err := u.userRepository.GetUserById(int(session.UserId))
	if err != nil {
		return nil, "", err
	}
	role, err := u.userRepository.GetDetailRole(user.RoleId)
	if err != nil {
		return nil, "", err
	}
	accessTokenStr, err := a.GenerateAccessToken(user.Username, role.RoleName)
	if err != nil {
		return nil, "", err
	}
	accessToken := dto.UserLoginResBody{
		AccessToken: accessTokenStr,
		Username:    user.Username,
		Email:       user.Email,
	}
	return &accessToken, accessTokenStr, err
}

func (u *authUsecaseImpl) Logout(refreshToken string) (*dto.UserLogoutResBody, error) {
	var err error
	a := utils.NewAuthUtil()
	_, err = a.ValidateToken(refreshToken, config.Config.AuthConfig.HmacSecretRefreshToken)
	if err != nil {
		return nil, err
	}

	session, err := u.authRepository.GetAuthSessionByRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	user, err := u.userRepository.GetUserById(int(session.UserId))
	if err != nil {
		return nil, err
	}

	err = u.authRepository.DeleteAuthSessionById(session.ID)
	if err != nil {
		return nil, err
	}

	res := dto.UserLogoutResBody{
		Username: user.Username,
		Email:    user.Email,
	}

	return &res, err
}
