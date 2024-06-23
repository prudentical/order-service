package service

import (
	"fmt"
	"order-service/internal/model"
	"reflect"
)

type FieldRequiredError struct {
	Field string
	Type  string
}

func (e FieldRequiredError) Error() string {
	return fmt.Sprintf("%s is required for %s", e.Field, e.Type)
}

type NotFoundError struct {
	Type interface{}
	Id   int64
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID[%d] not found!", reflect.TypeOf(n.Type).Name(), n.Id)
}

type InvalidOrderTypeError struct {
	Type model.OrderType
}

func (i InvalidOrderTypeError) Error() string {
	return fmt.Sprintf("Invalid order type[%s]", i.Type)
}

type ModifyingClosedPositionError struct {
	Id int64
}

func (m ModifyingClosedPositionError) Error() string {
	return fmt.Sprintf("The position ID[%d] is already closed", m.Id)
}

type InvalidUserAccount struct {
	UserId    int64
	AccountId int64
}

func (m InvalidUserAccount) Error() string {
	return fmt.Sprintf("User[id=%d] and Account[id=%d] are not associated", m.UserId, m.AccountId)
}
