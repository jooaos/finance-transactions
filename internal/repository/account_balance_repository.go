package repository

import "github.com/jooaos/pismo/internal/model"

type IAccountBalanceRepository interface {
	GetByAccountId(accountId int) (*model.AccountBalance, error)
	UpdateBalance(accountId int, balance float32) (*model.AccountBalance, error)
}
