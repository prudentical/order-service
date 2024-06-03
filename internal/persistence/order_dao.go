package persistence

import (
	"order-service/internal/model"

	"gorm.io/gorm"
)

type OrderDAO interface {
	Create(order model.Order) (model.Order, error)
}

type OrderDAOImpl struct {
	db *gorm.DB
}

func NewOrderDAO(conn *gorm.DB) OrderDAO {
	return OrderDAOImpl{conn}
}

func (dao OrderDAOImpl) Create(order model.Order) (model.Order, error) {
	tx := dao.db.Create(&order)
	if tx.Error != nil {
		return model.Order{}, tx.Error
	}
	return order, nil
}
