package day_5

import "testing"

func TestGetOverlapCount(t *testing.T) {
	tests := []struct {
		path        string
		expected    int
		description string
	}{
		{path: "partial_input.txt", expected: 5, description: "Partial input"},
		{path: "full_input.txt", expected: 5690, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetOverlapCount(test.path)

			if actual != test.expected {
				t.Errorf(
					"GetOverlapCount(%s) returned %v, expected %v",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}
