package usecase

import (
	"encoding/json"
	"sushee-backend/dto"
	"sushee-backend/entity"
	"sushee-backend/httperror/domain"
	"sushee-backend/repository"
)

type CartUsecase interface {
	GetCart(username string) (*[]dto.CartResBody, error)
	AddCart(username string, c *dto.CartReqBody) (*dto.CartResBody, error)
	DeleteCartsByUsername(username string) error
	DeleteCartByCartId(username string, cartId int) error
	UpdateCartByCartId(username string, cartId int, updatePremises *dto.CartEditDetailsReqBody) (*dto.CartResBody, error)
}

type cartUsecaseImpl struct {
	cartRepository      repository.CartRepository
	userRepository      repository.UserRepository
	menuRepository      repository.MenuRepository
	promotionRepository repository.PromotionRepository
}

type CartUsecaseConfig struct {
	CartRepository      repository.CartRepository
	UserRepository      repository.UserRepository
	MenuRepository      repository.MenuRepository
	PromotionRepository repository.PromotionRepository
}

func NewCartUsecase(c CartUsecaseConfig) CartUsecase {
	return &cartUsecaseImpl{
		cartRepository:      c.CartRepository,
		userRepository:      c.UserRepository,
		menuRepository:      c.MenuRepository,
		promotionRepository: c.PromotionRepository,
	}
}

func validateCartOwnership(cart *entity.Cart, user *entity.User) error {
	if cart.UserId != int(user.ID) {
		return domain.ErrCartUsecaseInvalidCartOwner
	}
	return nil
}

func (u *cartUsecaseImpl) GetCart(username string) (*[]dto.CartResBody, error) {
	carts, err := u.cartRepository.GetCartByUsername(username)
	if err != nil {
		return nil, err
	}

	newCarts := []dto.CartResBody{}
	for _, cart := range *carts {
		var customizations []entity.ChosenMenuOption
		err = json.Unmarshal([]byte(cart.MenuOption.Bytes), &customizations)
		if err != nil {
			return nil, domain.ErrCartUsecaseUnmarshallMenuOption
		}

		newCarts = append(newCarts, dto.CartResBody{
			ID:             cart.ID,
			UserId:         cart.UserId,
			MenuId:         cart.MenuId,
			PromotionId:    cart.PromotionId,
			Quantity:       cart.Quantity,
			MenuOption:     customizations,
			IsOrdered:      cart.IsOrdered,
			PromotionPrice: cart.PromotionPrice,
		})
	}
	return &newCarts, nil
}

func (u *cartUsecaseImpl) AddCart(username string, c *dto.CartReqBody) (*dto.CartResBody, error) {
	if c.PromotionId == nil && c.MenuId == nil {
		return nil, domain.ErrCartUsecaseMenuOrPromotionMenuNotFound
	}
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, err
	}

	newCartItem := entity.Cart{
		UserId:      int(user.ID),
		MenuId:      c.MenuId,
		PromotionId: c.PromotionId,
		Quantity:    c.Quantity,
		MenuOption:  c.MenuOption,
	}

	if _, err := json.Marshal(c.MenuOption); err != nil {
		return nil, domain.ErrCartUsecaseMarshallMenuOption
	}

	if c.PromotionId != nil {
		c, err := u.promotionRepository.GetAndValidatePromoMenu(*c.MenuId, *c.PromotionId)
		if err != nil {
			return nil, err
		}

		newCartItem.PromotionPrice = &c.PromotionPrice
	}

	createdCart, err := u.cartRepository.AddItemToCart(&newCartItem)
	if err != nil {
		return nil, err
	}

	var customizations []entity.ChosenMenuOption
	err = json.Unmarshal([]byte(createdCart.MenuOption.Bytes), &customizations)
	if err != nil {
		return nil, domain.ErrCartUsecaseUnmarshallMenuOption
	}

	cart := dto.CartResBody{
		ID:             createdCart.ID,
		UserId:         createdCart.UserId,
		MenuId:         createdCart.MenuId,
		PromotionId:    createdCart.PromotionId,
		Quantity:       createdCart.Quantity,
		MenuOption:     customizations,
		IsOrdered:      createdCart.IsOrdered,
		PromotionPrice: createdCart.PromotionPrice,
	}

	return &cart, nil
}

func (u *cartUsecaseImpl) DeleteCartsByUsername(username string) error {
	err := u.cartRepository.DeleteCartsByUsername(username)
	if err != nil {
		return err
	}

	return nil
}

func (u *cartUsecaseImpl) DeleteCartByCartId(username string, cartId int) error {
	cart, err := u.cartRepository.GetCartByCartId(cartId)
	if err != nil {
		return err
	}

	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return err
	}

	err = validateCartOwnership(cart, user)
	if err != nil {
		return err
	}

	err = u.cartRepository.DeleteCartByCartId(cartId)
	if err != nil {
		return err
	}

	return nil
}

func (u *cartUsecaseImpl) UpdateCartByCartId(username string, cartId int, reqBody *dto.CartEditDetailsReqBody) (*dto.CartResBody, error) {
	cart, err := u.cartRepository.GetCartByCartId(cartId)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, err
	}

	err = validateCartOwnership(cart, user)
	if err != nil {
		return nil, err
	}

	newCart := entity.Cart{
		Quantity: reqBody.Quantity,
	}

	err = u.cartRepository.UpdateCartByCartId(cartId, &newCart)
	if err != nil {
		return nil, err
	}

	updatedCart, err := u.cartRepository.GetCartByCartId(cartId)
	if err != nil {
		return nil, err
	}

	var customizations []entity.ChosenMenuOption
	err = json.Unmarshal([]byte(updatedCart.MenuOption.Bytes), &customizations)
	if err != nil {
		return nil, domain.ErrCartUsecaseUnmarshallMenuOption
	}

	resCart := dto.CartResBody{
		ID:             updatedCart.ID,
		UserId:         updatedCart.UserId,
		MenuId:         updatedCart.MenuId,
		PromotionId:    updatedCart.PromotionId,
		Quantity:       updatedCart.Quantity,
		MenuOption:     customizations,
		IsOrdered:      updatedCart.IsOrdered,
		PromotionPrice: updatedCart.PromotionPrice,
	}

	return &resCart, err
}
