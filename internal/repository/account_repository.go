package repository

import "github.com/jooaos/pismo/internal/model"

type IAccountResponsitory interface {
	Create(account *model.Account) (*model.Account, error)
	GetById(id uint32) (*model.Account, error)
	GetByDocumentNumber(documentNumber string) (*model.Account, error)
}
