package day_6

import "testing"

func TestGetLaternfishCount(t *testing.T) {
	tests := []struct {
		path        string
		expected    int
		description string
	}{
		{path: "partial_input.txt", expected: 5934, description: "Partial input"},
		{path: "full_input.txt", expected: 345793, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetLaternfishCount(test.path)

			if actual != test.expected {
				t.Errorf(
					"GetLaternfishCount(%s) returned %v, expected %v",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}
