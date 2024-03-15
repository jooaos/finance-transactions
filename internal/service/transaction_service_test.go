package service

import (
	"errors"
	"testing"

	"github.com/jooaos/pismo/internal/model"
	repository "github.com/jooaos/pismo/internal/tests/mocks/respository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransactionService_CreateTransaction_Success(t *testing.T) {
	transactionRespository := new(repository.TransactionRepositoryMock)
	accountRepository := new(repository.AccountRepositoryMock)
	accountBalanceRepository := new(repository.AccountBalanceMock)
	transactionService := NewTransactionService(transactionRespository, accountRepository, accountBalanceRepository)

	accountRepository.On("GetById", mock.Anything).Return(&model.Account{}, nil)
	accountBalanceRepository.On("UpdateBalance", mock.Anything, mock.Anything).Return(&model.AccountBalance{}, nil)

	var cases = []struct {
		accountId      int
		accountBalance float32
		operationType  int
		amount         float32
	}{
		{
			1,
			20,
			int(model.CASH_PURCHASE),
			-10.20,
		},
		{
			1,
			20,
			int(model.INSTALLMENT_PURCHASE),
			-11.20,
		},
		{
			1,
			20,
			int(model.WITHDRAWAL),
			-12.20,
		},
		{
			1,
			20,
			int(model.PAYMENT),
			13.20,
		},
	}

	for _, item := range cases {
		t.Run("should create transactions with success", func(t *testing.T) {
			accountBalanceRepository.On("GetByAccountId", mock.Anything).Return(model.NewAccountBalance(1, item.accountBalance), nil).Times(1)
			transaction := model.NewTransaction(item.accountId, item.operationType, item.amount)
			transactionRespository.On("Create", mock.Anything).Return(transaction, nil).Times(1)

			result, err := transactionService.CreateTransaction(item.accountId, item.operationType, item.amount)

			assert.Nil(t, err)
			assert.Equal(t, transaction, result)
		})
	}
}

func TestTransactionService_CreateTransaction_BalanceNotEnough(t *testing.T) {
	transactionRespository := new(repository.TransactionRepositoryMock)
	accountRepository := new(repository.AccountRepositoryMock)
	accountBalanceRepository := new(repository.AccountBalanceMock)
	transactionService := NewTransactionService(transactionRespository, accountRepository, accountBalanceRepository)

	accountRepository.On("GetById", mock.Anything).Return(&model.Account{}, nil)

	var cases = []struct {
		accountId      int
		accountBalance float32
		operationType  int
		amount         float32
	}{
		{
			1,
			5,
			int(model.CASH_PURCHASE),
			-10.20,
		},
		{
			1,
			5,
			int(model.INSTALLMENT_PURCHASE),
			-11.20,
		},
		{
			1,
			5,
			int(model.WITHDRAWAL),
			-12.20,
		},
	}

	for _, item := range cases {
		t.Run("should create transactions with success", func(t *testing.T) {
			accountBalanceRepository.On("GetByAccountId", mock.Anything).Return(model.NewAccountBalance(1, item.accountBalance), nil).Times(1)
			_, err := transactionService.CreateTransaction(item.accountId, item.operationType, item.amount)
			assert.NotNil(t, err)
		})
	}
}

func TestTransactionService_CreateTransaction_AmmoutNotAllowed(t *testing.T) {
	transactionRespository := new(repository.TransactionRepositoryMock)
	accountRepository := new(repository.AccountRepositoryMock)
	accountBalanceRepository := new(repository.AccountBalanceMock)
	transactionService := NewTransactionService(transactionRespository, accountRepository, accountBalanceRepository)

	var cases = []struct {
		accountId     int
		operationType int
		amount        float32
	}{
		{
			1,
			int(model.CASH_PURCHASE),
			10.20,
		},
		{
			1,
			int(model.INSTALLMENT_PURCHASE),
			11.20,
		},
		{
			1,
			int(model.WITHDRAWAL),
			12.20,
		},
		{
			1,
			int(model.PAYMENT),
			-13.20,
		},
	}

	for _, item := range cases {
		t.Run("should return amount not allowed error", func(t *testing.T) {
			_, err := transactionService.CreateTransaction(item.accountId, item.operationType, item.amount)
			assert.NotNil(t, err)
			assert.ErrorIs(t, err, ErrTransactionAmountNotAllowed)
		})
	}
}

func TestTransactionService_CreateTransaction_ValidationErros(t *testing.T) {
	transactionRespository := new(repository.TransactionRepositoryMock)
	accountRepository := new(repository.AccountRepositoryMock)
	accountBalanceRepository := new(repository.AccountBalanceMock)
	transactionService := NewTransactionService(transactionRespository, accountRepository, accountBalanceRepository)

	t.Run("should return operation type not found error", func(t *testing.T) {
		_, err := transactionService.CreateTransaction(2, 99999, 10)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrOperationTypeNotFound)
	})

	t.Run("should return account not found error", func(t *testing.T) {
		accountRepository.On("GetById", mock.Anything).Return(&model.Account{}, errors.New("database error"))
		_, err := transactionService.CreateTransaction(2, 4, 10)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrAccountNotFound)
	})
}
