package service

import (
	"errors"
	"testing"

	"github.com/jooaos/pismo/internal/model"
	repository "github.com/jooaos/pismo/internal/tests/mocks/respository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountService_Create(t *testing.T) {
	account := model.NewAccount("12345678900")
	accountRepository := new(repository.AccountRepositoryMock)
	accountRepository.On("Create", mock.Anything).Return(account, nil)
	accountRepository.On("GetByDocumentNumber", mock.Anything).Return(&model.Account{}, errors.New("test"))
	accountService := NewAccountService(accountRepository)
	t.Run("should create account", func(t *testing.T) {
		result, err := accountService.Create("12345678900")
		assert.Nil(t, err)
		assert.Equal(t, account, result)
	})

	t.Run("should return validate document error", func(t *testing.T) {
		_, err := accountService.Create("12345")
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrDocumentMustHave11Digits)
	})

	t.Run("should return account already exists error", func(t *testing.T) {
		accountRepository := new(repository.AccountRepositoryMock)
		accountRepository.On("GetByDocumentNumber", mock.Anything).Return(&model.Account{}, nil)
		accountService = NewAccountService(accountRepository)
		_, err := accountService.Create("12345678900")
		assert.ErrorIs(t, err, ErrAccountAlreadyExists)
	})
}

func TestAccountService_GetById(t *testing.T) {
	account := model.NewAccount("12345678900")
	account.ID = 1
	t.Run("should get by id", func(t *testing.T) {
		accountRepository := new(repository.AccountRepositoryMock)
		accountRepository.On("GetById", mock.Anything).Return(account, nil)
		accountService := NewAccountService(accountRepository)
		result, err := accountService.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, account, result)
	})

	t.Run("should return error from database", func(t *testing.T) {
		accountRepository := new(repository.AccountRepositoryMock)
		accountRepository.On("GetById", mock.Anything).Return(account, errors.New("any error"))
		accountService := NewAccountService(accountRepository)
		_, err := accountService.GetById(1)
		assert.NotNil(t, err)
	})
}
