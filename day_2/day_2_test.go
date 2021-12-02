package day_2

import "testing"

func TestGetFinalPositionProduct(t *testing.T) {
	tests := []struct {
		path        string
		expected    int
		description string
	}{
		{path: "partial_input.txt", expected: 150, description: "Partial input"},
		{path: "full_input.txt", expected: 2019945, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetFinalPositionProduct(test.path)

			if actual != test.expected {
				t.Errorf(
					"GetFinalPositionProduct(%s) returned %v, expected %v",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}
