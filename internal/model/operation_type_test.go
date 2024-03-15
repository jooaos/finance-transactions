package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperationType_IDExists(t *testing.T) {
	var cases = []struct {
		id     int
		result bool
	}{
		{
			1,
			true,
		},
		{
			2,
			true,
		},
		{
			3,
			true,
		},
		{
			4,
			true,
		},
		{
			5,
			false,
		},
	}

	for _, item := range cases {
		t.Run("should validate operation type id", func(t *testing.T) {
			result := ValidateOperationType(item.id)
			assert.Equal(t, item.result, result)
		})
	}
}

func TestOperationType_CorrectIDToOperation(t *testing.T) {
	var cases = []struct {
		operationType OperationTypeId
		result        int
	}{
		{
			CASH_PURCHASE,
			1,
		},
		{
			INSTALLMENT_PURCHASE,
			2,
		},
		{
			WITHDRAWAL,
			3,
		},
		{
			PAYMENT,
			4,
		},
	}

	for _, item := range cases {
		t.Run("should validate correct id to operation", func(t *testing.T) {
			assert.Equal(t, int(item.operationType), item.result)
		})
	}
}
