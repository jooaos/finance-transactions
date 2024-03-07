package repository

import "github.com/jooaos/pismo/internal/model"

type ITransactionRepository interface {
	Create(transaction *model.Transaction) (*model.Transaction, error)
}
