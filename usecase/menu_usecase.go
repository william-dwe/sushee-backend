package usecase

import (
	"encoding/json"
	"sushee-backend/dto"
	"sushee-backend/entity"
	"sushee-backend/httperror/domain"
	"sushee-backend/repository"
)

type MenuUsecase interface {
	GetMenu(q dto.MenuQuery) (*dto.MenusResBody, error)
	AddMenu(reqBody *dto.MenuAddReqBody) (*entity.Menu, error)
	UpdateMenuByMenuId(menuId int, m *entity.Menu) (*entity.Menu, error)
	DeleteMenuByMenuId(menuId int) error
	GetMenuDetailByMenuId(menuId int) (*entity.Menu, error)
}

type menuUsecaseImpl struct {
	menuRepository repository.MenuRepository
}

type MenuUsecaseConfig struct {
	MenuRepository repository.MenuRepository
}

func NewMenuUsecase(c MenuUsecaseConfig) MenuUsecase {
	return &menuUsecaseImpl{
		menuRepository: c.MenuRepository,
	}
}

func (u *menuUsecaseImpl) GetMenu(q dto.MenuQuery) (*dto.MenusResBody, error) {
	rows, err := u.menuRepository.GetMenuCount(q)
	if err != nil {
		return nil, err
	}

	if q.Limit == 0 {
		q.Limit = rows
	}
	menus, err := u.menuRepository.GetMenu(q)
	if err != nil {
		return nil, err
	}

	menuItems := []dto.MenuItem{}
	for _, menu := range *menus {
		var customizations []entity.MenuCustomization
		err = json.Unmarshal([]byte(menu.Customization.Bytes), &customizations)
		if err != nil {
			return nil, domain.ErrMenuUsecaseUnmarshallCustomizese
		}
		newMenu := dto.MenuItem{
			MenuName:          menu.MenuName,
			AvgRating:         menu.AvgRating,
			NumberOfFavorites: menu.NumberOfFavorites,
			Price:             menu.Price,
			MenuPhoto:         menu.MenuPhoto,
			CategoryId:        menu.CategoryId,
			Customization:     customizations,
		}
		menuItems = append(menuItems, newMenu)
	}

	maxPage := (rows + q.Limit - 1) / q.Limit
	resBody := dto.MenusResBody{
		Menus:       menuItems,
		CurrentPage: q.Page,
		MaxPage:     maxPage,
	}

	return &resBody, nil
}

func (u *menuUsecaseImpl) AddMenu(reqBody *dto.MenuAddReqBody) (*entity.Menu, error) {
	newMenu := entity.Menu{
		MenuName:      reqBody.MenuName,
		Price:         reqBody.Price,
		MenuPhoto:     reqBody.MenuPhoto,
		CategoryId:    reqBody.CategoryId,
		Customization: reqBody.Customization,
	}

	if _, err := json.Marshal(newMenu.Customization); err != nil {
		return nil, domain.ErrMenuUsecaseMarshallCustomizese
	}

	menu, err := u.menuRepository.AddMenu(&newMenu)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (u *menuUsecaseImpl) UpdateMenuByMenuId(menuId int, m *entity.Menu) (*entity.Menu, error) {
	err := u.menuRepository.UpdateMenuByMenuId(menuId, m)
	if err != nil {
		return nil, err
	}

	menu, err := u.menuRepository.GetMenuByMenuId(menuId)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (u *menuUsecaseImpl) DeleteMenuByMenuId(menuId int) error {
	err := u.menuRepository.DeleteMenuByMenuId(menuId)
	if err != nil {
		return err
	}

	return nil
}

func (u *menuUsecaseImpl) GetMenuDetailByMenuId(menuId int) (*entity.Menu, error) {
	menu, err := u.menuRepository.GetMenuDetailByMenuId(menuId)
	if err != nil {
		return nil, err
	}

	return menu, nil
}
