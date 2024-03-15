package adapter

import (
	"log"

	"github.com/jooaos/pismo/internal/model"
	"gorm.io/gorm"
)

type AccountBalanceRepositoryMariaDB struct {
	conn *gorm.DB
}

func (ac *AccountBalanceRepositoryMariaDB) GetByAccountId(accountId int) (*model.AccountBalance, error) {
	accountBalance := &model.AccountBalance{}

	sql := "SELECT * FROM accounts_balance WHERE account_id = ?"

	result := ac.conn.Raw(sql, accountId).First(&accountBalance)
	if result.Error != nil {
		log.Printf(
			"[AccountBalanceRepositoryMariaDB::GetByAccountId] Error while getting account balance: %s",
			result.Error.Error(),
		)
		return nil, result.Error
	}

	return accountBalance, nil
}
