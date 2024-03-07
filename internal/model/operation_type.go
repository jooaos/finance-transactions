package model

type OperationTypeId int

const (
	CASH_PURCHASE        OperationTypeId = 1
	INSTALLMENT_PURCHASE OperationTypeId = 2
	WITHDRAWAL           OperationTypeId = 3
	PAYMENT              OperationTypeId = 4
)

type OperationType struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func ValidateOperationType(operationType int) bool {
	for _, item := range getAll() {
		if int(item) == operationType {
			return true
		}
	}
	return false
}

func getAll() []OperationTypeId {
	return []OperationTypeId{
		CASH_PURCHASE,
		INSTALLMENT_PURCHASE,
		WITHDRAWAL,
		PAYMENT,
	}
}
