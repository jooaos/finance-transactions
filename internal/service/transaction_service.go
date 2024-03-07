package service

import (
	"github.com/jooaos/pismo/internal/model"
	"github.com/jooaos/pismo/internal/repository"
)

type TransactionService struct {
	transactionRepository repository.ITransactionRepository
	accountRepository     repository.IAccountResponsitory
}

type ITransactionService interface {
	CreateTransaction(accountId, operationTypeId int, amount float32) (*model.Transaction, error)
}

func NewTransactionService(
	transactionRepository repository.ITransactionRepository,
	accountRepository repository.IAccountResponsitory,
) *TransactionService {
	return &TransactionService{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
	}
}

func (tr *TransactionService) CreateTransaction(accountId, operationTypeId int, amount float32) (*model.Transaction, error) {
	_, err := tr.accountRepository.GetById(uint32(accountId))
	if err != nil {
		return nil, ErrAccountNotFound
	}

	chekOperationType := model.ValidateOperationType(operationTypeId)
	if !chekOperationType {
		return nil, ErrOperationTypeNotFound
	}

	checkAmount := model.ValidateAmount(operationTypeId, amount)
	if !checkAmount {
		return nil, ErrTransactionAmountNotAllowed
	}

	transaction := model.NewTransaction(accountId, operationTypeId, amount)

	result, err := tr.transactionRepository.Create(transaction)
	if err != nil {
		return nil, err
	}

	return result, nil
}
