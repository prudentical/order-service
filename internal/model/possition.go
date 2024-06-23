package model

type Position struct {
	BaseEntity
	BotId  int64          `json:"botId" validate:"required"`
	Status PositionStatus `json:"status" validate:"required"`
	Orders []Order        `json:"orders" gorm:"foreignKey:position_id"`
}

type PositionStatus string

const (
	Open   PositionStatus = "Open"
	Closed PositionStatus = "Closed"
)
