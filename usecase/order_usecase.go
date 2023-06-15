package usecase

import (
	"encoding/json"
	"sushee-backend/dto"
	"sushee-backend/entity"
	"sushee-backend/httperror/domain"
	"sushee-backend/repository"
	"time"
)

type OrderUsecase interface {
	AddOrder(username string, reqBody *dto.OrderReqBody) (*dto.OrderResBody, error)
	GetOrderStatus(oq *dto.OrderStatusQuery) (*dto.OrdersResBody, error)
	UpdateOrderStatus(reqBody *dto.OrderStatusUpdateReqBody) (*dto.OrderResBody, error)
	GetOrderHistory(username string, oq *dto.OrderHistoryQuery) (*dto.OrdersResBody, error)
}

type orderUsecaseImpl struct {
	orderRepository  repository.OrderRepository
	userRepository   repository.UserRepository
	cartRepository   repository.CartRepository
	couponRepository repository.CouponRepository
}

type OrderUsecaseConfig struct {
	OrderRepository  repository.OrderRepository
	UserRepository   repository.UserRepository
	CartRepository   repository.CartRepository
	CouponRepository repository.CouponRepository
}

func NewOrderUsecase(c OrderUsecaseConfig) OrderUsecase {
	return &orderUsecaseImpl{
		orderRepository:  c.OrderRepository,
		userRepository:   c.UserRepository,
		cartRepository:   c.CartRepository,
		couponRepository: c.CouponRepository,
	}
}

func validateCartOwnershipAndAvailability(cart *entity.Cart, user *entity.User) error {
	if cart.UserId != int(user.ID) {
		return domain.ErrOrderUsecaseUnauthorizedOrder
	}
	if cart.IsOrdered {
		return domain.ErrOrderUsecaseOrderIsOrdered
	}

	return nil
}

func (u *orderUsecaseImpl) AddOrder(username string, reqBody *dto.OrderReqBody) (*dto.OrderResBody, error) {
	if len(reqBody.CartIdList) == 0 {
		return nil, domain.ErrOrderUsecaseOrderEmpty
	}
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, err
	}

	carts, err := u.cartRepository.GetCartByCartIds(reqBody.CartIdList)
	if err != nil {
		return nil, err
	}
	if len(*carts) != len(reqBody.CartIdList) {
		return nil, domain.ErrOrderUsecaseSomeCartIdNotFound
	}
	for _, cart := range *carts {
		err = validateCartOwnershipAndAvailability(&cart, user)
		if err != nil {
			return nil, err
		}
	}

	newOrder := entity.Order{
		UserId:          int(user.ID),
		OrderDate:       time.Now(),
		PaymentOptionId: reqBody.PaymentOptionId,
		Status:          entity.ORDER_STATUS_PAYMENT,
	}

	var couponId int
	if reqBody.CouponCode != "" {
		coupon, err := u.couponRepository.GetUserCouponByCouponCode(int(user.ID), reqBody.CouponCode)
		if err != nil {
			return nil, err
		}
		couponId = int(coupon.ID)
		newOrder.CouponId = &couponId
		newOrder.DiscountAmount = coupon.DiscountAmount
	}

	totalPrice, err := u.cartRepository.GetCartTotalPriceByCartIds(reqBody.CartIdList)
	if err != nil {
		return nil, err
	}
	newOrder.GrossAmount = totalPrice
	newOrder.NetAmount = newOrder.GrossAmount - newOrder.DiscountAmount

	var newOrderedMenuResBodys []dto.OrderedMenuResBody
	var newOrderedMenus []entity.OrderedMenu
	for _, cart := range *carts {
		orderedMenu := entity.OrderedMenu{
			Quantity:   cart.Quantity,
			MenuOption: cart.MenuOption,
		}
		if cart.PromotionId != nil {
			orderedMenu.PromotionId = cart.PromotionId
		}
		if cart.MenuId != nil {
			orderedMenu.MenuId = cart.MenuId
		}

		newOrderedMenus = append(newOrderedMenus, orderedMenu)

		orderedMenuResBody := dto.OrderedMenuResBody{
			Quantity: cart.Quantity,
		}
		if cart.MenuId != nil {
			orderedMenuResBody.MenuId = *cart.MenuId
		}
		if cart.PromotionId != nil {
			orderedMenuResBody.PromotionId = *cart.PromotionId
		}

		var customizations []entity.ChosenMenuOption
		err = json.Unmarshal([]byte(cart.MenuOption.Bytes), &customizations)
		if err != nil {
			return nil, domain.ErrCartUsecaseUnmarshallMenuOption
		}
		orderedMenuResBody.Customization = customizations
		newOrderedMenuResBodys = append(newOrderedMenuResBodys, orderedMenuResBody)
	}

	newOrder.OrderedMenus = newOrderedMenus

	order, err := u.orderRepository.AddOrder(&newOrder)
	if err != nil {
		return nil, err
	}

	err = u.cartRepository.UpdateCartByCartIds(reqBody.CartIdList, &entity.Cart{
		IsOrdered: true,
	})
	if err != nil {
		return nil, err
	}

	res := dto.OrderResBody{
		ID:              int(order.ID),
		UserId:          order.UserId,
		OrderDate:       order.OrderDate,
		PaymentOptionId: order.PaymentOptionId,
		OrderedMenus:    newOrderedMenuResBodys,
		GrossAmount:     order.GrossAmount,
		DiscountAmount:  order.DiscountAmount,
		NetAmount:       order.NetAmount,
		Status:          order.Status,
	}

	if reqBody.CouponCode != "" {
		res.CouponId = *newOrder.CouponId
		u.couponRepository.DeleteCouponById(*newOrder.CouponId)
	}

	return &res, nil
}

func (u *orderUsecaseImpl) GetOrderStatus(oq *dto.OrderStatusQuery) (*dto.OrdersResBody, error) {
	orders, err := u.orderRepository.GetOrderStatus(*oq)
	if err != nil {
		return nil, err
	}

	oSlice := []dto.OrderResBody{}
	for _, o := range *orders {
		var newOrderedMenuResBodys []dto.OrderedMenuResBody
		for _, om := range o.OrderedMenus {
			orderedMenuResBody := dto.OrderedMenuResBody{
				Quantity: om.Quantity,
			}
			if om.MenuId != nil {
				orderedMenuResBody.MenuId = *om.MenuId
			}
			if om.PromotionId != nil {
				orderedMenuResBody.PromotionId = *om.PromotionId
			}
			var customizations []entity.ChosenMenuOption
			err = json.Unmarshal([]byte(om.MenuOption.Bytes), &customizations)
			if err != nil {
				return nil, domain.ErrCartUsecaseUnmarshallMenuOption
			}
			orderedMenuResBody.Customization = customizations
			newOrderedMenuResBodys = append(newOrderedMenuResBodys, orderedMenuResBody)
		}

		oSlice = append(oSlice, dto.OrderResBody{
			ID:              int(o.ID),
			UserId:          o.UserId,
			OrderDate:       o.OrderDate,
			CouponId:        int(*o.CouponId),
			PaymentOptionId: o.PaymentOptionId,
			OrderedMenus:    newOrderedMenuResBodys,
			GrossAmount:     o.GrossAmount,
			DiscountAmount:  o.DiscountAmount,
			NetAmount:       o.NetAmount,
			Status:          o.Status,
		})
	}

	res := dto.OrdersResBody{
		Orders: oSlice,
	}

	return &res, nil
}

func (u *orderUsecaseImpl) UpdateOrderStatus(reqBody *dto.OrderStatusUpdateReqBody) (*dto.OrderResBody, error) {
	orderWithNewStatus := entity.Order{
		Status: reqBody.Status,
	}

	err := u.orderRepository.UpdateOrderByOrderId(reqBody.ID, &orderWithNewStatus)
	if err != nil {
		return nil, err
	}

	order, err := u.orderRepository.GetOrderById(reqBody.ID)
	if err != nil {
		return nil, err
	}

	var newOrderedMenuResBodys []dto.OrderedMenuResBody
	for _, om := range order.OrderedMenus {
		orderedMenuResBody := dto.OrderedMenuResBody{
			Quantity: om.Quantity,
		}
		if om.MenuId != nil {
			orderedMenuResBody.MenuId = *om.MenuId
		}
		if om.PromotionId != nil {
			orderedMenuResBody.PromotionId = *om.PromotionId
		}
		var customizations []entity.ChosenMenuOption
		err = json.Unmarshal([]byte(om.MenuOption.Bytes), &customizations)
		if err != nil {
			return nil, domain.ErrCartUsecaseUnmarshallMenuOption
		}
		orderedMenuResBody.Customization = customizations
		newOrderedMenuResBodys = append(newOrderedMenuResBodys, orderedMenuResBody)
	}

	res := dto.OrderResBody{
		ID:              int(order.ID),
		UserId:          order.UserId,
		OrderDate:       order.OrderDate,
		CouponId:        int(*order.CouponId),
		PaymentOptionId: order.PaymentOptionId,
		OrderedMenus:    newOrderedMenuResBodys,
		GrossAmount:     order.GrossAmount,
		DiscountAmount:  order.DiscountAmount,
		NetAmount:       order.NetAmount,
		Status:          order.Status,
	}

	return &res, nil
}

func (u *orderUsecaseImpl) GetOrderHistory(username string, oq *dto.OrderHistoryQuery) (*dto.OrdersResBody, error) {
	user, err := u.userRepository.GetUserByEmailOrUsername(username)
	if err != nil {
		return nil, err
	}

	rows, err := u.orderRepository.GetOrderHistoryCount(int(user.ID), oq)
	if err != nil {
		return nil, err
	}

	orders, err := u.orderRepository.GetOrderHistory(int(user.ID), *oq)
	if err != nil {
		return nil, err
	}

	oSlice := []dto.OrderResBody{}
	for _, o := range *orders {
		var newOrderedMenuResBodys []dto.OrderedMenuResBody
		for _, om := range o.OrderedMenus {
			orderedMenuResBody := dto.OrderedMenuResBody{
				Quantity: om.Quantity,
			}
			if om.MenuId != nil {
				orderedMenuResBody.MenuId = *om.MenuId
			}
			if om.PromotionId != nil {
				orderedMenuResBody.PromotionId = *om.PromotionId
			}
			var customizations []entity.ChosenMenuOption
			err = json.Unmarshal([]byte(om.MenuOption.Bytes), &customizations)
			if err != nil {
				return nil, domain.ErrCartUsecaseUnmarshallMenuOption
			}
			orderedMenuResBody.Customization = customizations
			newOrderedMenuResBodys = append(newOrderedMenuResBodys, orderedMenuResBody)
		}

		no := dto.OrderResBody{
			ID:              int(o.ID),
			UserId:          o.UserId,
			OrderDate:       o.OrderDate,
			PaymentOptionId: o.PaymentOptionId,
			OrderedMenus:    newOrderedMenuResBodys,
			GrossAmount:     o.GrossAmount,
			DiscountAmount:  o.DiscountAmount,
			NetAmount:       o.NetAmount,
			Status:          o.Status,
		}
		if o.CouponId != nil {
			no.CouponId = *o.CouponId
		}

		oSlice = append(oSlice, no)
	}

	maxPage := (rows + oq.Limit - 1) / oq.Limit
	res := dto.OrdersResBody{
		Orders:      oSlice,
		CurrentPage: oq.Page,
		MaxPage:     maxPage,
	}

	return &res, nil
}
