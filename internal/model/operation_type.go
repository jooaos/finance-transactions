package model

type OperationTypeId int

const (
	A_VISTA          OperationTypeId = 1
	COMPRA_PARCELADA OperationTypeId = 2
	SAQUE            OperationTypeId = 3
	PAGAMENTO        OperationTypeId = 4
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
		A_VISTA,
		COMPRA_PARCELADA,
		SAQUE,
		PAGAMENTO,
	}
}
