package service

import (
	"order-service/internal/model"
	"order-service/internal/persistence"
)

type OrderService interface {
	GetByBotId(botId int, page int, size int) (persistence.Page[model.Order], error)
	DeleteByBotId(botId int) error
	Create(order model.Order) (model.Order, error)
}

type OrderServiceImpl struct {
	dao persistence.OrderDAO
}

func NewOrderService(dao persistence.OrderDAO) OrderService {
	return OrderServiceImpl{dao}
}

func (s OrderServiceImpl) Create(order model.Order) (model.Order, error) {
	order.ID = 0
	return s.dao.Create(order)
}

func (s OrderServiceImpl) GetByBotId(botId int, page int, size int) (persistence.Page[model.Order], error) {
	return s.dao.GetByBotIdPaginated(botId, page, size)
}

func (s OrderServiceImpl) DeleteByBotId(id int) error {
	return s.dao.DeleteByBotId(id)
}
