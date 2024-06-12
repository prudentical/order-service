package service

import (
	"order-service/internal/model"
	"order-service/internal/persistence"
)

type PositionService interface {
	GetByBotId(userId int64, accountId int64, botId int64, page int, size int, status string) (persistence.Page[model.Position], error)
	DeleteByBotId(userId int64, accountId int64, botId int64) error
	Create(position model.Position) (model.Position, error)
	Close(id int64) error
	GetById(id int64) (model.Position, error)
}

type positionServiceImpl struct {
	dao      persistence.PositionDAO
	accounts AccountService
}

func NewPositionService(dao persistence.PositionDAO, accounts AccountService) PositionService {
	return positionServiceImpl{dao, accounts}
}

func (s positionServiceImpl) Create(position model.Position) (model.Position, error) {
	position.ID = 0
	return s.dao.Create(position)
}
func (s positionServiceImpl) GetById(id int64) (model.Position, error) {
	return s.dao.GetById(id)
}

func (s positionServiceImpl) GetByBotId(userId int64, accountId int64, botId int64, page int, size int, status string) (persistence.Page[model.Position], error) {
	_, err := s.accounts.GetAccount(userId, accountId)
	if err != nil {
		return persistence.Page[model.Position]{}, err
	}
	var statusPtr *string
	if len(status) > 0 {
		statusPtr = &status
	}
	return s.dao.GetByBotIdPaginated(botId, page, size, statusPtr)
}

func (s positionServiceImpl) DeleteByBotId(userId int64, accountId int64, id int64) error {
	_, err := s.accounts.GetAccount(userId, accountId)
	if err != nil {
		return err
	}
	return s.dao.DeleteByBotId(id)
}

func (s positionServiceImpl) Close(id int64) error {
	position, err := s.dao.GetById(id)
	if err != nil {
		return err
	}
	position.Status = model.Closed

	return s.dao.Update(position)
}
