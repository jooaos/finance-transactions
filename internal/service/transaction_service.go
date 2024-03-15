package service

import (
	"log"
	"math"

	"github.com/jooaos/pismo/internal/model"
	"github.com/jooaos/pismo/internal/repository"
)

type TransactionService struct {
	transactionRepository    repository.ITransactionRepository
	accountRepository        repository.IAccountResponsitory
	accountBalanceRepository repository.IAccountBalanceRepository
}

type ITransactionService interface {
	CreateTransaction(accountId, operationTypeId int, amount float32) (*model.Transaction, error)
}

func NewTransactionService(
	transactionRepository repository.ITransactionRepository,
	accountRepository repository.IAccountResponsitory,
	accountBalanceRepository repository.IAccountBalanceRepository,
) *TransactionService {
	return &TransactionService{
		transactionRepository:    transactionRepository,
		accountRepository:        accountRepository,
		accountBalanceRepository: accountBalanceRepository,
	}
}

func (tr *TransactionService) CreateTransaction(accountId, operationTypeId int, amount float32) (*model.Transaction, error) {
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

	_, err := tr.accountRepository.GetById(uint32(accountId))
	if err != nil {
		log.Printf("[TransactionService::CreateTransaction] Account not found")
		return nil, ErrAccountNotFound
	}

	accountBalance, err := tr.accountBalanceRepository.GetByAccountId(accountId)
	if err != nil {
		return nil, ErrAccountBalanceNotFound
	}

	if operationTypeId != int(model.PAYMENT) && accountBalance.Balance < float32(math.Abs(float64(amount))) {
		return nil, ErrAccountBalanceNotEnough
	}

	// TODO: colocar dentro de uma transaction
	result, err := tr.addTransaction(model.NewTransaction(accountId, operationTypeId, amount))
	if err != nil {
		log.Printf("[TransactionService::CreateTransaction] Error while creating account: %s", err.Error())
		return nil, err
	}

	_, err = tr.updateBalance(accountId, accountBalance.Balance+amount)
	if err != nil {
		log.Printf("[TransactionService::CreateTransaction] Error while updating account balance: %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (tr *TransactionService) addTransaction(transaction *model.Transaction) (*model.Transaction, error) {
	result, err := tr.transactionRepository.Create(transaction)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (tr *TransactionService) updateBalance(accountId int, balance float32) (*model.AccountBalance, error) {
	result, err := tr.accountBalanceRepository.UpdateBalance(accountId, balance)
	if err != nil {
		return nil, err
	}
	return result, nil
}
