package repository

import (
	"sushee-backend/entity"
	"sushee-backend/httperror/domain"
	"sushee-backend/utils"

	"gorm.io/gorm"
)

type AuthRepository interface {
	AddAuthSession(s *entity.AuthSession) (*entity.AuthSession, error)
	GetAuthSessionByRefreshToken(t string) (*entity.AuthSession, error)
	DeleteAuthSessionById(id uint) error
}

type authRepositoryImpl struct {
	db *gorm.DB
}

type AuthRepositoryConfig struct {
	DB *gorm.DB
}

func NewAuthRepository(c AuthRepositoryConfig) AuthRepository {
	return &authRepositoryImpl{
		db: c.DB,
	}
}

func (r *authRepositoryImpl) AddAuthSession(s *entity.AuthSession) (*entity.AuthSession, error) {
	err := r.db.Create(&s).Error
	if err != nil {
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"auth_sesssions_user_id_fkey": domain.ErrAuthRepoAddAuthSessionUserNotFound,
			},
			domain.ErrAuthRepoInternal,
		)
		return nil, err
	}
	return s, err
}

func (r *authRepositoryImpl) GetAuthSessionByRefreshToken(t string) (*entity.AuthSession, error) {
	var s entity.AuthSession
	err := r.db.
		Where("refresh_token = ?", t).
		First(&s).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrAuthRepoAuthSessionNotFound
		}
		return nil, domain.ErrAuthRepoInternal
	}
	return &s, err
}

func (r *authRepositoryImpl) DeleteAuthSessionById(id uint) error {
	err := r.db.Delete(&id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.ErrAuthRepoAddAuthSessionUserNotFound
		}

		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"auth_sesssions_user_id_fkey": domain.ErrAuthRepoAddAuthSessionUserNotFound,
			},
			domain.ErrAuthRepoInternal,
		)
		return err
	}
	return nil
}
