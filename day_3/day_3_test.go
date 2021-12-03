package day_3

import "testing"

func TestGetPowerConsumption(t *testing.T) {
	tests := []struct {
		path        string
		expected    int64
		description string
	}{
		{path: "partial_input.txt", expected: 198, description: "Partial input"},
		{path: "full_input.txt", expected: 2743844, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetPowerConsumption(test.path)

			if actual != test.expected {
				t.Errorf(
					"GetPowerConsumption(%s) returned %v, expected %v",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}

func TestGetLifeSupportRating(t *testing.T) {
	tests := []struct {
		path        string
		expected    int64
		description string
	}{
		{path: "partial_input.txt", expected: 230, description: "Partial input"},
		{path: "full_input.txt", expected: 6677951, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetLifeSupportRating(test.path)

			if actual != test.expected {
				t.Errorf(
					"GetLifeSupportRating(%s) returned %v, expected %v",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}
