package service

import (
	"log"

	"github.com/jooaos/pismo/internal/model"
	"github.com/jooaos/pismo/internal/repository"
)

type AccountBalanceService struct {
	accountBalanceRespository repository.IAccountBalanceRepository
}

type IAccountBalanceService interface {
	GetByAccountId(accountId int) (*model.AccountBalance, error)
	UpdateBalance(accountId int, balance float32) (*model.AccountBalance, error)
}

func NewAccountBalanceService(accountBalanceRepository repository.IAccountBalanceRepository) *AccountBalanceService {
	return &AccountBalanceService{
		accountBalanceRespository: accountBalanceRepository,
	}
}

func (ac *AccountBalanceService) GetByAccountId(accountId int) (*model.AccountBalance, error) {
	result, err := ac.accountBalanceRespository.GetByAccountId(accountId)
	if err != nil {
		log.Printf("[AccountBalanceService::GetByAccountId] Error while getting account balance: %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (ac *AccountBalanceService) UpdateBalance(accountId int, balance float32) (*model.AccountBalance, error) {
	result, err := ac.accountBalanceRespository.UpdateBalance(accountId, balance)
	if err != nil {
		log.Printf("[AccountBalanceService::UpdateBalance] Error while updating account balance: %s", err.Error())
		return nil, err
	}

	return result, nil
}
