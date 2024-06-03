package service

import (
	"errors"
	"order-service/internal/dto"
	"order-service/internal/model"
	"order-service/internal/persistence"

	"github.com/shopspring/decimal"
)

type OrderService interface {
	Create(order dto.OrderDTO) (model.Order, error)
}

type orderServiceImpl struct {
	dao       persistence.OrderDAO
	positions PositionService
}

func NewOrderService(dao persistence.OrderDAO, position PositionService) OrderService {
	return orderServiceImpl{dao, position}
}

func (s orderServiceImpl) Create(order dto.OrderDTO) (model.Order, error) {
	switch order.Type {
	case model.Buy:
		position, err := s.positions.Create(model.Position{BotId: order.BotId, Status: model.Open})
		if err != nil {
			return model.Order{}, err
		}
		order.PositionId = position.ID
	case model.Sell:
		if order.PositionId == 0 {
			return model.Order{}, FieldRequiredError{Field: "PositionId", Type: "sell orders"}
		}
		position, err := s.positions.GetById(order.PositionId)
		if err != nil {
			if errors.Is(err, persistence.RecordNotFoundError{}) {
				return model.Order{}, NotFoundError{Type: model.Position{}, Id: order.PositionId}
			}
		}
		if position.Status == model.Closed {
			return model.Order{}, ModifyingClosedPositionError{position.ID}
		}
		amount := calcOpenAmount(position)
		if amount.Equal(order.FilledAmount) {
			err := s.positions.Close(position.ID)
			if err != nil {
				return model.Order{}, err
			}
		}
	default:
		return model.Order{}, InvalidOrderTypeError{order.Type}
	}
	return s.dao.Create(order.ToOrder())
}

func calcOpenAmount(position model.Position) decimal.Decimal {
	sum := decimal.NewFromInt32(0)
	for _, order := range position.Orders {
		switch order.Type {
		case model.Buy:
			sum = sum.Add(order.FilledAmount)
		case model.Sell:
			sum = sum.Sub(order.FilledAmount)
		}
	}
	return sum
}
