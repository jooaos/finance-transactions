package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateOperationType(t *testing.T) {
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
