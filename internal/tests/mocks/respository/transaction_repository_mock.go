package repository

import (
	"github.com/jooaos/pismo/internal/model"
	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (ac *TransactionRepositoryMock) Create(transaction *model.Transaction) (*model.Transaction, error) {
	args := ac.Called(transaction)
	return args.Get(0).(*model.Transaction), args.Error(1)
}
