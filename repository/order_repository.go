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

type OrderRepository interface {
	AddOrder(newOrder *entity.Order) (*entity.Order, error)
	AddOrderedMenu(newOrderedMenus *[]entity.OrderedMenu) (*[]entity.OrderedMenu, error)
	GetOrderStatus(oq dto.OrderStatusQuery) (*[]entity.Order, error)
	GetOrderHistoryCount(userId int, oq *dto.OrderHistoryQuery) (int, error)
	GetOrderHistory(userId int, oq dto.OrderHistoryQuery) (*[]entity.Order, error)
	GetOrderById(orderId int) (*entity.Order, error)
	UpdateOrderByOrderId(orderId int, newOrderStatus *entity.Order) error
	GetOrderedMenuById(orderedMenuId int) (*entity.OrderedMenu, error)
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

type OrderRepositoryConfig struct {
	DB *gorm.DB
}

func NewOrderRepository(c OrderRepositoryConfig) OrderRepository {
	return &orderRepositoryImpl{
		db: c.DB,
	}
}

func (r *orderRepositoryImpl) AddOrder(newOrder *entity.Order) (*entity.Order, error) {
	err := r.db.
		Create(newOrder).
		Error
	if err != nil {
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"orders_coupon_id_fkey":         domain.ErrOrderRepoCouponNotFound,
				"orders_payment_option_id_fkey": domain.ErrOrderRepoPaymentOptionNotFound,
				"orders_user_id_fkey":           domain.ErrOrderRepoUserNotFound,
			},
			domain.ErrOrderRepoInternal,
		)
		return nil, err
	}
	return newOrder, nil
}

func (r *orderRepositoryImpl) AddOrderedMenu(newOrderedMenus *[]entity.OrderedMenu) (*[]entity.OrderedMenu, error) {
	err := r.db.
		Create(newOrderedMenus).
		Error
	if err != nil {
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"ordered_menus_menu_id_fkey":      domain.ErrOrderRepoMenuNotFound,
				"ordered_menus_order_id_fkey":     domain.ErrOrderRepoOrderNotFound,
				"ordered_menus_promotion_id_fkey": domain.ErrOrderRepoPromotionNotFound,
			},
			domain.ErrOrderRepoInternal,
		)
		return nil, err
	}
	return newOrderedMenus, nil
}

func (r *orderRepositoryImpl) GetOrderStatus(oq dto.OrderStatusQuery) (*[]entity.Order, error) {
	var o []entity.Order

	sqSelectedMenuOrder := r.db.
		Table("orders o").
		Select("user_id, order_id, menu_id").
		Joins("join ordered_menus om ON o.id = om.order_id").
		Where("status in (?)", oq.FilterByStatus)

	sqSelectedMenu := r.db.
		Table("menus m").
		Select("order_id").
		Joins("join (?) sm  ON sm.menu_id = m.id", sqSelectedMenuOrder).
		Where("menu_name ilike (?)", "%"+oq.Search+"%")

	q := r.db.
		Preload("OrderedMenus").
		Preload("OrderedMenus.Menu").
		Preload("OrderedMenus.Review").
		Where("id in (?)", sqSelectedMenu).
		Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: oq.SortBy,
			},
			Desc: strings.ToLower(oq.Sort) == "desc",
		}).
		Limit(oq.Limit).
		Offset(oq.Page*oq.Limit - oq.Limit).
		Find(&o)
	err := q.Error
	if err != nil {
		return nil, domain.ErrOrderRepoInternal
	}
	return &o, nil
}

func (r *orderRepositoryImpl) GetOrderHistoryCount(userId int, oq *dto.OrderHistoryQuery) (int, error) {
	var rows int64

	sqSelectedMenuOrder := r.db.
		Table("orders o").
		Select("user_id, order_id, menu_id").
		Joins("join ordered_menus om ON o.id = om.order_id").
		Where("user_id in (?)", userId)
	sqSelectedMenu := r.db.
		Table("menus m").
		Select("order_id").
		Joins("join (?) sm  ON sm.menu_id = m.id", sqSelectedMenuOrder).
		Where("menu_name ilike (?)", "%"+oq.Search+"%")

	q := r.db.
		Where("id in (?)", sqSelectedMenu).
		Table("orders")
	q.Count(&rows)
	err := q.Error
	if err != nil {
		return 0, domain.ErrOrderRepoInternal
	}
	return int(rows), nil
}

func (r *orderRepositoryImpl) GetOrderHistory(userId int, oq dto.OrderHistoryQuery) (*[]entity.Order, error) {
	var o []entity.Order

	sqSelectedMenuOrder := r.db.
		Table("orders o").
		Select("user_id, order_id, menu_id").
		Joins("join ordered_menus om ON o.id = om.order_id").
		Where("user_id in (?)", userId)

	sqSelectedMenu := r.db.
		Table("menus m").
		Select("order_id").
		Joins("join (?) sm  ON sm.menu_id = m.id", sqSelectedMenuOrder).
		Where("menu_name ilike (?)", "%"+oq.Search+"%")

	q := r.db.
		Preload("OrderedMenus").
		Preload("OrderedMenus.Menu").
		Preload("OrderedMenus.Review").
		Where("id in (?)", sqSelectedMenu).
		Order(clause.OrderByColumn{
			Column: clause.Column{
				Name: oq.SortBy,
			},
			Desc: strings.ToLower(oq.Sort) == "desc",
		}).
		Limit(oq.Limit).
		Offset(oq.Page*oq.Limit - oq.Limit).
		Find(&o)
	err := q.Error
	if err != nil {
		return nil, domain.ErrOrderRepoInternal
	}
	return &o, nil
}

func (r *orderRepositoryImpl) GetOrderById(orderId int) (*entity.Order, error) {
	var o entity.Order

	err := r.db.
		Where("id = (?)", orderId).
		First(&o).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrOrderRepoOrderNotFound
		}
		return nil, domain.ErrOrderRepoInternal
	}
	return &o, nil
}

func (r *orderRepositoryImpl) UpdateOrderByOrderId(orderId int, newOrderStatus *entity.Order) error {
	err := r.db.
		Where("id = ?", orderId).
		Updates(newOrderStatus).
		Error
	if err != nil {
		err = utils.PgConsErrMasker(
			err,
			entity.ConstraintErrMaskerMap{
				"orders_coupon_id_fkey":         domain.ErrOrderRepoCouponNotFound,
				"orders_payment_option_id_fkey": domain.ErrOrderRepoPaymentOptionNotFound,
				"orders_user_id_fkey":           domain.ErrOrderRepoUserNotFound,
			},
			domain.ErrOrderRepoInternal,
		)
		return err
	}
	return nil
}

func (r *orderRepositoryImpl) GetOrderedMenuById(orderedMenuId int) (*entity.OrderedMenu, error) {
	var o entity.OrderedMenu

	err := r.db.
		Where("id = (?)", orderedMenuId).
		First(&o).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrOrderRepoOrderNotFound
		}
		return nil, domain.ErrOrderRepoInternal
	}
	return &o, err
}
