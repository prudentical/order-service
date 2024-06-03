package service

import (
	"order-service/internal/model"
	"order-service/internal/persistence"
)

type PositionService interface {
	GetByBotId(botId int64, page int, size int, status string) (persistence.Page[model.Position], error)
	DeleteByBotId(botId int64) error
	Create(position model.Position) (model.Position, error)
	Close(id int64) error
	GetById(id int64) (model.Position, error)
}

type positionServiceImpl struct {
	dao persistence.PositionDAO
}

func NewPositionService(dao persistence.PositionDAO) PositionService {
	return positionServiceImpl{dao}
}

func (s positionServiceImpl) Create(position model.Position) (model.Position, error) {
	position.ID = 0
	return s.dao.Create(position)
}
func (s positionServiceImpl) GetById(id int64) (model.Position, error) {
	return s.dao.GetById(id)
}

func (s positionServiceImpl) GetByBotId(botId int64, page int, size int, status string) (persistence.Page[model.Position], error) {
	var statusPtr *string
	if len(status) > 0 {
		statusPtr = &status
	}
	return s.dao.GetByBotIdPaginated(botId, page, size, statusPtr)
}

func (s positionServiceImpl) DeleteByBotId(id int64) error {
	return s.dao.DeleteByBotId(id)
}

func (s positionServiceImpl) Close(id int64) error {
	position, err := s.GetById(id)
	if err != nil {
		return err
	}
	position.Status = model.Closed

	return s.dao.Update(position)
}
