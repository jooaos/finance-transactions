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
			CASH_PURCHASE,
			-10.0,
			true,
		},
		{
			CASH_PURCHASE,
			10.0,
			false,
		},
		{
			INSTALLMENT_PURCHASE,
			-10.0,
			true,
		},
		{
			INSTALLMENT_PURCHASE,
			10.0,
			false,
		},
		{
			WITHDRAWAL,
			-10.0,
			true,
		},
		{
			WITHDRAWAL,
			10.0,
			false,
		},
		{
			PAYMENT,
			-10.0,
			false,
		},
		{
			PAYMENT,
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
