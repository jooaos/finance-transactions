package adapter

import (
	"fmt"
	"log"

	"github.com/jooaos/pismo/internal/model"
	"gorm.io/gorm"
)

type AccountRepositoryMariaDB struct {
	conn *gorm.DB
}

func NewAccountRepositoryMariaDB(connection *gorm.DB) *AccountRepositoryMariaDB {
	return &AccountRepositoryMariaDB{
		conn: connection,
	}
}

func (ac *AccountRepositoryMariaDB) Create(account *model.Account) (*model.Account, error) {
	sql := "INSERT INTO accounts (document_number) VALUES (?) RETURNING *"

	result := ac.conn.Raw(sql, account.DocumentNumber).Scan(&account)
	if result.Error != nil {
		fmt.Printf(
			"[AccountRepositoryMariaDB::Create] Error while creating account: %s",
			result.Error.Error(),
		)
		return nil, result.Error
	}

	return account, nil
}

func (ac *AccountRepositoryMariaDB) GetById(id uint32) (*model.Account, error) {
	account := &model.Account{}

	sql := "SELECT * FROM accounts WHERE id = ?"

	result := ac.conn.Raw(sql, id).First(&account)
	if result.Error != nil {
		log.Printf(
			"[AccountRepositoryMariaDB::GetById] Error while getting account: %s",
			result.Error.Error(),
		)
		return nil, result.Error
	}

	return account, nil
}

func (ac *AccountRepositoryMariaDB) GetByDocumentNumber(documentNumber string) (*model.Account, error) {
	account := &model.Account{}

	sql := "SELECT * FROM accounts WHERE document_number = ?"

	result := ac.conn.Raw(sql, documentNumber).First(&account)
	if result.Error != nil {
		log.Printf(
			"[AccountRepositoryMariaDB::GetByDocumentNumber] Error while getting account: %s",
			result.Error.Error(),
		)
		return nil, result.Error
	}

	return account, nil
}
