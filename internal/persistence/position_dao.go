package persistence

import (
	"errors"
	"order-service/internal/model"

	"gorm.io/gorm"
)

type PositionDAO interface {
	Create(position model.Position) (model.Position, error)
	GetById(id int64) (model.Position, error)
	GetByBotIdPaginated(botId int64, page int, size int, status *string) (Page[model.Position], error)
	DeleteByBotId(botId int64) error
	Update(position model.Position) error
}

type positionDAOImpl struct {
	db *gorm.DB
}

func NewPositionDAO(conn *gorm.DB) PositionDAO {
	return positionDAOImpl{conn}
}

func (dao positionDAOImpl) Create(position model.Position) (model.Position, error) {
	tx := dao.db.Create(&position)
	if tx.Error != nil {
		return model.Position{}, tx.Error
	}
	return position, nil
}
func (dao positionDAOImpl) GetById(id int64) (model.Position, error) {
	var position model.Position
	tx := dao.db.Preload("Orders").Find(&position, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return model.Position{}, RecordNotFoundError{}
		}
		return model.Position{}, tx.Error
	}
	return position, nil
}

func (dao positionDAOImpl) GetByBotIdPaginated(botId int64, page int, size int, status *string) (Page[model.Position], error) {
	var positions []model.Position = make([]model.Position, 1)
	query := dao.db.Scopes(Paginate(page, size))
	query = OptionalFilter(query, "status", status)

	tx := query.Preload("Orders").Find(&positions, "bot_id = ?", botId)
	if tx.Error != nil {
		return Page[model.Position]{}, tx.Error
	}
	var total int64
	query = dao.db.Model(model.Position{}).Where("bot_id = ?", botId)
	query = OptionalFilter(query, "status", status)
	query.Count(&total)
	return Page[model.Position]{
		List:  positions,
		Page:  page,
		Size:  len(positions),
		Total: total,
	}, nil
}

func (dao positionDAOImpl) GetByPositionId(positionId int64) ([]model.Position, error) {
	var positions []model.Position = make([]model.Position, 1)

	tx := dao.db.Preload("Orders").Find(&positions, "position_id = ?", positionId)
	if tx.Error != nil {
		return []model.Position{}, tx.Error
	}
	return positions, nil
}

func (dao positionDAOImpl) DeleteByBotId(botId int64) error {
	tx := dao.db.Delete(model.Position{}, "bot_id = ?", botId)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (dao positionDAOImpl) Update(position model.Position) error {
	tx := dao.db.Save(&position)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
