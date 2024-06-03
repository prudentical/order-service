package dto

import (
	"order-service/internal/model"
	"time"

	"github.com/shopspring/decimal"
)

type OrderDTO struct {
	PositionId   int64             `json:"positionId"`
	BotId        int64             `json:"botId"  validate:"required"`
	InternalId   string            `json:"internalId" validate:"required"`
	Amount       decimal.Decimal   `json:"amount" validate:"required"`
	FilledAmount decimal.Decimal   `json:"filledAmount" validate:"required"`
	Price        decimal.Decimal   `json:"price" validate:"required"`
	Type         model.OrderType   `json:"type" validate:"required"`
	Status       model.OrderStatus `json:"status" validate:"required"`
	DateTime     time.Time         `json:"datetime" validate:"required"`
}

func (o OrderDTO) ToOrder() model.Order {
	return model.Order{
		PositionId:   o.PositionId,
		InternalId:   o.InternalId,
		Amount:       o.Amount,
		FilledAmount: o.FilledAmount,
		Price:        o.Price,
		Type:         o.Type,
		Status:       o.Status,
		DateTime:     o.DateTime,
	}
}
