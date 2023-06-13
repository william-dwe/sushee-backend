package usecase

import (
	"sushee-backend/utils"

	"sushee-backend/dto"
	"sushee-backend/entity"
	"sushee-backend/repository"
)

type UserUsecase interface {
	GetDetailUserByUsername(accessToken string) (*dto.UserContext, error)
	UpdateUserDetailsByUsername(username string, updatePremises dto.UserEditDetailsReqBody) (*dto.UserContext, error)
	GetDetailRole(roleId int) (*entity.Role, error)
}

type userUsecaseImpl struct {
	userRepository repository.UserRepository
}

type UserUsecaseConfig struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(c UserUsecaseConfig) UserUsecase {
	return &userUsecaseImpl{
		userRepository: c.UserRepository,
	}
}

func (u *userUsecaseImpl) GetDetailUserByUsername(username string) (*dto.UserContext, error) {
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, err
	}

	userContext := dto.UserContext{
		Username:       user.Username,
		FullName:       user.FullName,
		Email:          user.Email,
		Phone:          user.Phone,
		ProfilePicture: user.ProfilePicture,
		PlayAttempt:    user.PlayAttempt,
		RoleId:         user.RoleId,
	}

	return &userContext, nil
}

func (u *userUsecaseImpl) UpdateUserDetailsByUsername(username string, reqBody dto.UserEditDetailsReqBody) (*dto.UserContext, error) {
	var err error

	hashedPass := ""
	if reqBody.Password != "" {
		hashedPass, _ = utils.HashAndSalt(reqBody.Password)
	}

	// TODO: upload image + embed image here
	// var uploadUrl string
	// var err error
	// if formFile.Img != nil {
	// 	uploadUrl, err = utils.NewMediaUpload().FileUpload(username, formFile)
	// 	if err != nil {
	// 		_ = c.Error(err)
	// 		return
	// 	}
	// }

	newUser := entity.User{
		FullName:       reqBody.FullName,
		Phone:          reqBody.Phone,
		Email:          reqBody.Email,
		Password:       hashedPass,
		ProfilePicture: reqBody.ProfilePicture,
	}

	user, err := u.userRepository.UpdateUserDetailsByUsername(username, &newUser)
	if err != nil {
		return nil, err
	}

	userContext := dto.UserContext{
		Username:       user.Username,
		FullName:       user.FullName,
		Email:          user.Email,
		Phone:          user.Phone,
		ProfilePicture: user.ProfilePicture,
		PlayAttempt:    user.PlayAttempt,
		RoleId:         user.RoleId,
	}
	return &userContext, err
}

func (u *userUsecaseImpl) GetDetailRole(roleId int) (*entity.Role, error) {
	role, err := u.userRepository.GetDetailRole(roleId)
	if err != nil {
		return nil, err
	}
	return role, nil
}
