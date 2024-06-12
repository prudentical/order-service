package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order-service/internal/discovery"
	"order-service/internal/dto"
)

type AccountService interface {
	GetAccount(userId int64, accountId int64) (dto.AccountDTO, error)
}

type accountServiceImpl struct {
	client           http.Client
	serviceDiscovery discovery.ServiceDiscovery
}

func NewAccountService(client http.Client, serviceDiscovery discovery.ServiceDiscovery) AccountService {
	return accountServiceImpl{client: client, serviceDiscovery: serviceDiscovery}
}

func (a accountServiceImpl) GetAccount(userId int64, accountId int64) (dto.AccountDTO, error) {
	serviceUrl, err := a.serviceDiscovery.Discover("account-service")
	if err != nil {
		return dto.AccountDTO{}, err
	}
	path := fmt.Sprintf("/users/%d/accounts/%d", userId, accountId)
	resp, err := a.client.Get(fmt.Sprintf("http://%s%s", serviceUrl, path))
	if err != nil {
		return dto.AccountDTO{}, err
	}
	if resp.StatusCode == 404 {
		return dto.AccountDTO{}, InvalidUserAccount{UserId: userId, AccountId: accountId}
	}
	var account dto.AccountDTO
	err = json.NewDecoder(resp.Body).Decode(&account)
	if err != nil {
		return dto.AccountDTO{}, err
	}
	return account, nil
}
