package service

import (
	"errors"
	"testing"

	"github.com/jooaos/pismo/internal/model"
	repository "github.com/jooaos/pismo/internal/tests/mocks/respository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountBalanceService_GetByAccountId(t *testing.T) {
	accountBalanceRepository := new(repository.AccountBalanceMock)
	accountBalanceService := NewAccountBalanceService(accountBalanceRepository)

	t.Run("should return account balance with success", func(t *testing.T) {
		accountBalanceRepository.On("GetByAccountId", mock.Anything).Return(model.NewAccountBalance(1, 1), nil).Times(1)
		result, err := accountBalanceService.GetByAccountId(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result.AccountId)
	})

	t.Run("should return return error", func(t *testing.T) {
		accountBalanceRepository.On("GetByAccountId", mock.Anything).Return(new(model.AccountBalance), errors.New("teste")).Times(1)
		_, err := accountBalanceService.GetByAccountId(1)
		assert.NotNil(t, err)
	})
}

func TestAccountBalanceService_UpdateBalance(t *testing.T) {
	accountBalanceRepository := new(repository.AccountBalanceMock)
	accountBalanceService := NewAccountBalanceService(accountBalanceRepository)

	t.Run("should return account balance updated ", func(t *testing.T) {
		accountBalanceRepository.On("UpdateBalance", 1, float32(20)).Return(model.NewAccountBalance(1, 1), nil).Times(1)
		result, err := accountBalanceService.UpdateBalance(1, 20)
		assert.Nil(t, err)
		assert.Equal(t, 1, result.AccountId)
	})

	t.Run("should return return error", func(t *testing.T) {
		accountBalanceRepository.On("UpdateBalance", 1, float32(20)).Return(new(model.AccountBalance), errors.New("teste")).Times(1)
		_, err := accountBalanceService.UpdateBalance(1, 20)
		assert.NotNil(t, err)
	})
}
