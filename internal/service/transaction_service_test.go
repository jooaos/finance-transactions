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
	accountRepository.On("GetById", mock.Anything).Return(&model.Account{}, nil)

	var cases = []struct {
		accountId     int
		operationType int
		amount        float32
	}{
		{
			1,
			int(model.A_VISTA),
			-10.20,
		},
		{
			1,
			int(model.COMPRA_PARCELADA),
			-11.20,
		},
		{
			1,
			int(model.SAQUE),
			-12.20,
		},
		{
			1,
			int(model.PAGAMENTO),
			13.20,
		},
	}

	for _, item := range cases {
		t.Run("should create transactions with success", func(t *testing.T) {
			transaction := model.NewTransaction(item.accountId, item.operationType, item.amount)
			transactionRespository.On("Create", mock.Anything).Return(transaction, nil).Times(1)
			transactionService := NewTransactionService(transactionRespository, accountRepository)
			result, err := transactionService.CreateTransaction(item.accountId, item.operationType, item.amount)
			assert.Nil(t, err)
			assert.Equal(t, transaction, result)
		})
	}

	var cases2 = []struct {
		accountId     int
		operationType int
		amount        float32
	}{
		{
			1,
			int(model.A_VISTA),
			10.20,
		},
		{
			1,
			int(model.COMPRA_PARCELADA),
			11.20,
		},
		{
			1,
			int(model.SAQUE),
			12.20,
		},
		{
			1,
			int(model.PAGAMENTO),
			-13.20,
		},
	}

	for _, item := range cases2 {
		t.Run("should return amount not allowed error", func(t *testing.T) {
			transactionService := NewTransactionService(transactionRespository, accountRepository)
			_, err := transactionService.CreateTransaction(item.accountId, item.operationType, item.amount)
			assert.NotNil(t, err)
			assert.ErrorIs(t, err, ErrTransactionAmountNotAllowed)
		})
	}

	t.Run("should return operation type not found error", func(t *testing.T) {
		transactionService := NewTransactionService(transactionRespository, accountRepository)
		_, err := transactionService.CreateTransaction(2, 99999, 10)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrOperationTypeNotFound)
	})

	t.Run("should return operation type not found error", func(t *testing.T) {
		accountRepository := new(repository.AccountRepositoryMock)
		accountRepository.On("GetById", mock.Anything).Return(&model.Account{}, errors.New("database error"))
		transactionService := NewTransactionService(transactionRespository, accountRepository)
		_, err := transactionService.CreateTransaction(2, 99999, 10)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, ErrAccountNotFound)
	})
}
