package repository

import (
	"sushee-backend/entity"
	"sushee-backend/httperror/domain"
	"sushee-backend/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(i int) (*entity.User, error)
	GetUserByEmailOrUsername(string) (*entity.User, error)
	AddNewUser(*entity.User) (*entity.User, error)
	UpdateUserDetailsByUsername(username string, newUser *entity.User) (res *entity.User, lastErr error)
	GetDetailRole(roleId int) (*entity.Role, error)
	CheckDuplicatePhone(phone string) (bool, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

type UserRepositoryConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c UserRepositoryConfig) UserRepository {
	return &userRepositoryImpl{
		db: c.DB,
	}
}

func (r *userRepositoryImpl) GetUserById(i int) (*entity.User, error) {
	var user entity.User
	err := r.db.
		Where("id = ?", i).
		First(&user).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserRepoUserNotFound
		}
		return nil, domain.ErrUserRepoInternal
	}
	return &user, err
}

func (r *userRepositoryImpl) GetUserByEmailOrUsername(i string) (*entity.User, error) {
	var user entity.User
	err := r.db.
		Where("email = ?", i).
		Or("username = ?", i).
		First(&user).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserRepoUserNotFound
		}
		return nil, domain.ErrUserRepoInternal
	}

	return &user, err
}

func (r *userRepositoryImpl) AddNewUser(u *entity.User) (*entity.User, error) {
	err := r.db.
		Create(&u).
		Error
	if err != nil {
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"users_role_id_fkey":   domain.ErrUserRepoRoleNotFound,
				"users_email_key":      domain.ErrUserRepoEmailAlreadyExist,
				"users_phone_key":      domain.ErrUserRepoPhoneAlreadyExist,
				"users_phone_check":    domain.ErrUserRepoInvalidPhoneFormat,
				"users_username_key":   domain.ErrUserRepoUsernameAlreadyExist,
				"users_username_check": domain.ErrUserRepoInvalidUsernameFormat,
			},
			domain.ErrAuthRepoInternal,
		)
	}
	return u, err
}

func (r *userRepositoryImpl) AlterUserDetailsByUsernameTx(tx *gorm.DB, username string, newUser *entity.User) error {
	var user entity.User
	err := tx.Model(&user).
		Where("username = ?", username).
		Updates(newUser).
		Debug().
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.ErrUserRepoUserNotFound
		}
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"users_role_id_fkey":   domain.ErrUserRepoRoleNotFound,
				"users_email_key":      domain.ErrUserRepoEmailAlreadyExist,
				"users_phone_key":      domain.ErrUserRepoPhoneAlreadyExist,
				"users_phone_check":    domain.ErrUserRepoInvalidPhoneFormat,
				"users_username_key":   domain.ErrUserRepoUsernameAlreadyExist,
				"users_username_check": domain.ErrUserRepoInvalidUsernameFormat,
			},
			domain.ErrAuthRepoInternal,
		)
	}
	return err
}

func (r *userRepositoryImpl) RetrieveUserDetailsByUsernameTx(tx *gorm.DB, u string) (*entity.User, error) {
	var user entity.User
	err := tx.
		Or("username = ?", u).
		First(&user).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserRepoUserNotFound
		}
		return nil, domain.ErrUserRepoInternal
	}

	return &user, err
}

func (r *userRepositoryImpl) UpdateUserDetailsByUsername(username string, newUser *entity.User) (res *entity.User, lastErr error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			lastErr = domain.ErrAuthRepoInternal
		}
	}()

	err := r.AlterUserDetailsByUsernameTx(tx, username, newUser)
	if err != nil {
		return nil, err
	}

	user, err := r.RetrieveUserDetailsByUsernameTx(tx, username)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, domain.ErrUserRepoInternal
	}
	return user, nil
}

func (r *userRepositoryImpl) GetDetailRole(roleId int) (*entity.Role, error) {
	var role entity.Role
	err := r.db.
		Where("id = ?", roleId).
		First(&role).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrUserRepoDetailRoleNotFound
		}
		return nil, domain.ErrUserRepoInternal
	}
	return &role, err
}

func (r *userRepositoryImpl) CheckDuplicatePhone(phone string) (bool, error) {
	var user entity.User
	err := r.db.
		Where("phone = ? AND phone != ''", phone).
		First(&user).
		Error

	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return true, domain.ErrUserRepoInternal
	}
	return true, domain.ErrUserRepoPhoneAlreadyExist
}
