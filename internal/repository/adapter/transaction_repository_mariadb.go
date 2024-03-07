package adapter

import (
	"log"

	"github.com/jooaos/pismo/internal/model"
	"gorm.io/gorm"
)

type TransactionRepositoryMariaDB struct {
	conn *gorm.DB
}

func NewTransactionRepositoryMariaDB(connection *gorm.DB) *TransactionRepositoryMariaDB {
	return &TransactionRepositoryMariaDB{
		conn: connection,
	}
}

func (tr *TransactionRepositoryMariaDB) Create(transaction *model.Transaction) (*model.Transaction, error) {
	sql := "INSERT INTO transactions (account_id, operation_type_id, amount) VALUES (?,?,?)"

	result := tr.conn.Raw(
		sql,
		transaction.AccountId,
		transaction.OperationTypeId,
		transaction.Amount,
	).Scan(&transaction)
	if result.Error != nil {
		log.Printf(
			"[TransactionRepositoryMariaDB::Create] Error while creating transaction: %s",
			result.Error.Error(),
		)
		return nil, result.Error
	}

	return transaction, nil
}
