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
}

func NewAccountBalanceService(accountBalanceRepository repository.IAccountBalanceRepository) *AccountBalanceService {
	return &AccountBalanceService{
		accountBalanceRespository: accountBalanceRepository,
	}
}

func (ac *AccountBalanceService) GetByAccountId(accountId int) (*model.AccountBalance, error) {
	result, err := ac.accountBalanceRespository.GetByAccountId(accountId)
	if err != nil {
		log.Printf("[AccountBalanceService::GetByAccountId] Document number is not valid")
		return nil, err
	}

	return result, nil
}
