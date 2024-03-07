package service

import (
	"log"

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
		log.Printf("[TransactionService::CreateTransaction] Account not found")
		return nil, ErrAccountNotFound
	}

	chekOperationType := model.ValidateOperationType(operationTypeId)
	if !chekOperationType {
		log.Printf("[TransactionService::CreateTransaction] Operation type is not correct")
		return nil, ErrOperationTypeNotFound
	}

	checkAmount := model.ValidateAmount(operationTypeId, amount)
	if !checkAmount {
		log.Printf("[TransactionService::CreateTransaction] Amount is not correct")
		return nil, ErrTransactionAmountNotAllowed
	}

	transaction := model.NewTransaction(accountId, operationTypeId, amount)

	result, err := tr.transactionRepository.Create(transaction)
	if err != nil {
		log.Printf("[TransactionService::CreateTransaction] Error while creating account: %s", err.Error())
		return nil, err
	}

	return result, nil
}
