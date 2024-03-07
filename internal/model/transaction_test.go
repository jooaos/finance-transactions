package model

import "testing"

func TestValidateAmount(t *testing.T) {
	var cases = []struct {
		operationType OperationTypeId
		amount        float32
		result        bool
	}{
		{
			A_VISTA,
			-10.0,
			true,
		},
		{
			A_VISTA,
			10.0,
			false,
		},
		{
			COMPRA_PARCELADA,
			-10.0,
			true,
		},
		{
			COMPRA_PARCELADA,
			10.0,
			false,
		},
		{
			SAQUE,
			-10.0,
			true,
		},
		{
			SAQUE,
			10.0,
			false,
		},
		{
			PAGAMENTO,
			-10.0,
			false,
		},
		{
			PAGAMENTO,
			10.0,
			true,
		},
	}

	for _, item := range cases {
		t.Run("ValidateAmount", func(t *testing.T) {
			result := ValidateAmount(int(item.operationType), item.amount)
			if result != item.result {
				t.Errorf("Expected for operation type %d response be %t and got %t", item.operationType, item.result, result)
			}
		})
	}
}
