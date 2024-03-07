package service

import (
	"log"

	"github.com/jooaos/pismo/internal/model"
	"github.com/jooaos/pismo/internal/repository"
)

type AccountService struct {
	accountRepository repository.IAccountResponsitory
}

type IAccountService interface {
	Create(documentNumber string) (*model.Account, error)
	GetById(id uint32) (*model.Account, error)
}

func NewAccountService(accountRepository repository.IAccountResponsitory) *AccountService {
	return &AccountService{
		accountRepository: accountRepository,
	}
}

func (ac *AccountService) Create(documentNumber string) (*model.Account, error) {
	validate := ac.validateDocument(documentNumber)
	if !validate {
		log.Printf("[AccountService::Create] Document number is not valid")
		return nil, ErrDocumentMustHave11Digits
	}

	_, err := ac.accountRepository.GetByDocumentNumber(documentNumber)
	if err == nil {
		return nil, ErrAccountAlreadyExists
	}

	result, err := ac.accountRepository.Create(model.NewAccount(documentNumber))
	if err != nil {
		log.Printf("[AccountService::Create] Error while creating account: %s", err.Error())
		return nil, err
	}

	return result, nil
}

func (ac *AccountService) GetById(id uint32) (*model.Account, error) {
	account, err := ac.accountRepository.GetById(id)
	if err == nil {
		return account, nil
	}

	log.Printf("[AccountService::GetById] Error while getting account: %s", err.Error())
	return nil, err
}

func (ac AccountService) validateDocument(documentNumber string) bool {
	return len(documentNumber) == 11
}
