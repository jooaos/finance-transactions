package model

type AccountBalance struct {
	AccountId int     `json:"account_id"`
	Balance   float32 `json:"balance"`
}

func NewAccountBalance(accountId int, balance float32) *AccountBalance {
	return &AccountBalance{
		AccountId: accountId,
		Balance:   balance,
	}
}
