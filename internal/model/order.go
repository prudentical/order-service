package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	BaseEntity
	PositionId   int64           `json:"positionId"`
	InternalId   string          `json:"internalId" validate:"required"`
	Amount       decimal.Decimal `json:"amount" validate:"required"`
	FilledAmount decimal.Decimal `json:"filledAmount" validate:"required"`
	Price        decimal.Decimal `json:"price" validate:"required"`
	Type         OrderType       `json:"type" validate:"required"`
	Status       OrderStatus     `json:"status" validate:"required"`
	DateTime     time.Time       `json:"datetime" validate:"required"`
}

type OrderType string

const (
	Buy  OrderType = "Buy"
	Sell OrderType = "Sell"
)

type OrderStatus string

const (
	Pending   OrderStatus = "Pending"
	Partial   OrderStatus = "Partial"
	Fulfilled OrderStatus = "Fulfilled"
	Canceled  OrderStatus = "Canceled"
)
