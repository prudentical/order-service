package persistence

import (
	"order-service/internal/model"

	"gorm.io/gorm"
)

type OrderDAO interface {
	Create(order model.Order) (model.Order, error)
	GetByBotIdPaginated(botId int, page int, size int) (Page[model.Order], error)
	DeleteByBotId(botId int) error
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

func (dao OrderDAOImpl) GetByBotIdPaginated(botId int, page int, size int) (Page[model.Order], error) {
	var orders []model.Order = make([]model.Order, 1)
	tx := dao.db.Scopes(Paginate(page, size)).Find(&orders, "bot_id = ?", botId)
	if tx.Error != nil {
		return Page[model.Order]{}, tx.Error
	}
	var total int64
	dao.db.Model(model.Order{}).Where("bot_id = ?", botId).Count(&total)
	return Page[model.Order]{
		List:  orders,
		Page:  page,
		Size:  len(orders),
		Total: total,
	}, nil
}

func (dao OrderDAOImpl) DeleteByBotId(botId int) error {
	tx := dao.db.Delete(model.Order{}, "bot_id = ?", botId)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
