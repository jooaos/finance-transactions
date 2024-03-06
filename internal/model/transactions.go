package model

type Transactions struct {
	ID              int     `json:"id"`
	AccountId       int     `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float32 `json:"amount"`
}
