package model

type Transactions struct {
	ID              int             `json:"id"`
	AccountId       int             `json:"account_id"`
	OperationTypeId OperationTypeId `json:"operation_type_id"`
	Amount          float32         `json:"amount"`
}
