package models

import "testing"

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
		t.Run("ValidateOperationType", func(t *testing.T) {
			result := ValidateOperationType(item.id)
			if result != item.result {
				t.Errorf("Expected for id %d response be %t and got %t", item.id, item.result, result)
			}
		})
	}
}
