package dto

import "github.com/shopspring/decimal"

type AccountDTO struct {
	ID            int64           `json:"id"`
	Name          string          `json:"name"`
	UserId        int64           `json:"userId"`
	ExchangeId    int64           `json:"exchangeId"`
	Capital       decimal.Decimal `json:"capital"`
	LockedCapital decimal.Decimal `json:"lockedCapital"`
}
