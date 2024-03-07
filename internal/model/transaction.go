package model

type Transaction struct {
	ID              int             `json:"id"`
	AccountId       int             `json:"account_id"`
	OperationTypeId OperationTypeId `json:"operation_type_id"`
	Amount          float32         `json:"amount"`
}

func NewTransaction(accountId, operationTypeId int, amount float32) *Transaction {
	return &Transaction{
		AccountId:       accountId,
		OperationTypeId: OperationTypeId(operationTypeId),
		Amount:          amount,
	}
}

func ValidateAmount(operationTypeId int, amount float32) bool {
	switch operationTypeId {
	case int(A_VISTA), int(COMPRA_PARCELADA), int(SAQUE):
		if amount >= 0 {
			return false
		}
	case int(PAGAMENTO):
		if amount <= 0 {
			return false
		}
	}
	return true
}
