package day_8

import "testing"

func TestGetUniqueDigits(t *testing.T) {
	tests := []struct {
		path        string
		expected    int
		description string
	}{
		{path: "partial_input.txt", expected: 26, description: "Partial input"},
		{path: "full_input.txt", expected: 452, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetUniqueDigits(test.path)

			if actual != test.expected {
				t.Errorf(
					"GetUniqueDigits(%s) returned %v, expected %v",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}

}
