package repository

import (
	"github.com/jooaos/pismo/internal/model"
	"github.com/stretchr/testify/mock"
)

type AccountBalanceMock struct {
	mock.Mock
}

func (ac *AccountBalanceMock) GetByAccountId(accountId int) (*model.AccountBalance, error) {
	args := ac.Called(accountId)
	return args.Get(0).(*model.AccountBalance), args.Error(1)
}
