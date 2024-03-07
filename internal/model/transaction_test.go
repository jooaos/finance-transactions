package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		t.Run("should validate amount of transaction", func(t *testing.T) {
			result := ValidateAmount(int(item.operationType), item.amount)
			assert.Equal(t, result, item.result)
		})
	}
}
