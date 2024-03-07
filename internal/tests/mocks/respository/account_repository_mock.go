package repository

import (
	"github.com/jooaos/pismo/internal/model"
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (ac *AccountRepositoryMock) Create(account *model.Account) (*model.Account, error) {
	args := ac.Called(account)
	return args.Get(0).(*model.Account), args.Error(1)
}

func (ac *AccountRepositoryMock) GetById(id uint32) (*model.Account, error) {
	args := ac.Called(id)
	return args.Get(0).(*model.Account), args.Error(1)
}

func (ac *AccountRepositoryMock) GetByDocumentNumber(documentNumber string) (*model.Account, error) {
	args := ac.Called(documentNumber)
	return args.Get(0).(*model.Account), args.Error(1)
}
