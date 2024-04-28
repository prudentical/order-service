package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	BaseEntity
	BotId    int             `json:"botId" validate:"required"`
	Amount   decimal.Decimal `json:"amount" validate:"required"`
	Price    decimal.Decimal `json:"price" validate:"required"`
	Type     OrderType       `json:"type" validate:"required"`
	DateTime time.Time       `json:"datetime" validate:"required"`
}

type OrderType string

const (
	Buy  OrderType = "Buy"
	Sell OrderType = "Sell"
)
