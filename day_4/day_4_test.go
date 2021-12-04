package day_4

import "testing"

func TestGetBingoScore(t *testing.T) {
	tests := []struct {
		path        string
		expected    int
		description string
	}{
		{path: "partial_input.txt", expected: 4512, description: "Partial input"},
		{path: "full_input.txt", expected: 8442, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetBingoScore(test.path)

			if actual != test.expected {
				t.Errorf(
					"GetBingoScore(%s) returned %v, expected %v",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}
