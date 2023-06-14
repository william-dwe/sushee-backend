package repository

import (
	"strings"
	"sushee-backend/dto"
	"sushee-backend/entity"
	"sushee-backend/httperror/domain"
	"sushee-backend/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MenuRepository interface {
	GetMenu(dto.MenuQuery) (*[]entity.Menu, error)
	GetMenuCount(q dto.MenuQuery) (int, error)
	AddMenu(newMenu *entity.Menu) (*entity.Menu, error)
	GetMenuByMenuId(menuId int) (*entity.Menu, error)
	UpdateMenuByMenuId(menuId int, newMenu *entity.Menu) error
	DeleteMenuByMenuId(menuId int) error
	GetMenuDetailByMenuId(menuId int) (*entity.Menu, error)
}

type menuRepositoryImpl struct {
	db *gorm.DB
}

type MenuRepositoryConfig struct {
	DB *gorm.DB
}

func NewMenuRepository(c MenuRepositoryConfig) MenuRepository {
	return &menuRepositoryImpl{
		db: c.DB,
	}
}

func (r *menuRepositoryImpl) GetMenuCount(q dto.MenuQuery) (int, error) {
	var rows int64

	menuCategorySQ := r.db.
		Select("id").
		Where("category_name IN (?)", strings.Split(q.FilterByCategory, ",")).
		Or("'' = ?", q.FilterByCategory).
		Table("categories")
	query := r.db.
		Joins("menus").
		Where("category_id IN (?)", menuCategorySQ).
		Table("menus")
	query.Count(&rows)
	err := query.Error
	if err != nil {
		return 0, domain.ErrMenuRepoInternal
	}
	return int(rows), nil
}

func (r *menuRepositoryImpl) GetMenu(q dto.MenuQuery) (*[]entity.Menu, error) {
	var menus []entity.Menu

	menuCategorySQ := r.db.
		Select("id").
		Where("category_name IN (?)", strings.Split(q.FilterByCategory, ",")).
		Or("'' = ?", q.FilterByCategory).
		Table("categories")
	menuSQ := r.db.
		Joins("menus").
		Where("category_id IN (?)", menuCategorySQ).
		Table("menus")
	query := r.db.
		Table("(?) as th", menuSQ).
		Where("menu_name ilike ?", "%"+q.Search+"%").
		Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: q.SortBy,
			},
			Desc: strings.ToLower(q.Sort) == "desc",
		}).
		Limit(q.Limit).
		Offset(q.Page*q.Limit - q.Limit).
		Find(&menus)
	err := query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrMenuRepoNoMenuFound
		}

		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"menus_category_id_fkey": domain.ErrMenuRepoCategoryNotFound,
			},
			domain.ErrMenuRepoInternal,
		)
		return nil, err
	}
	return &menus, nil
}

func (r *menuRepositoryImpl) AddMenu(newMenu *entity.Menu) (*entity.Menu, error) {
	err := r.db.
		Create(newMenu).
		Error
	if err != nil {
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"menus_category_id_fkey": domain.ErrMenuRepoCategoryNotFound,
			},
			domain.ErrMenuRepoInternal,
		)
		return nil, err
	}
	return newMenu, nil
}

func (r *menuRepositoryImpl) GetMenuByMenuId(menuId int) (*entity.Menu, error) {
	var menu entity.Menu

	err := r.db.
		Where("id = ?", menuId).
		First(&menu).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrMenuRepoMenuNotFound
		}
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"menus_category_id_fkey": domain.ErrMenuRepoCategoryNotFound,
			},
			domain.ErrMenuRepoInternal,
		)
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepositoryImpl) UpdateMenuByMenuId(menuId int, newMenu *entity.Menu) error {
	err := r.db.
		Where("id = ?", menuId).
		Updates(newMenu).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.ErrMenuRepoMenuNotFound
		}
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"menus_category_id_fkey": domain.ErrMenuRepoCategoryNotFound,
			},
			domain.ErrMenuRepoInternal,
		)
		return err
	}
	return nil
}

func (r *menuRepositoryImpl) DeleteMenuByMenuId(menuId int) error {
	var menu entity.Menu

	query := r.db.
		Where("id = (?)", menuId).
		Delete(&menu)

	err := query.Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.ErrMenuRepoMenuNotFound
		}
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"menus_category_id_fkey": domain.ErrMenuRepoCategoryNotFound,
			},
			domain.ErrMenuRepoInternal,
		)
		return err
	}

	return nil
}

func (r *menuRepositoryImpl) GetMenuDetailByMenuId(menuId int) (*entity.Menu, error) {
	var m entity.Menu

	q := r.db.
		Preload("Reviews").
		Where("id in (?)", menuId).
		Find(&m)
	err := q.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrMenuRepoMenuNotFound
		}
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"menus_category_id_fkey": domain.ErrMenuRepoCategoryNotFound,
			},
			domain.ErrMenuRepoInternal,
		)
		return nil, err
	}
	return &m, q.Error
}
