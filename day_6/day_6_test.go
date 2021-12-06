package day_6

import "testing"

func TestGetLaternfishCount(t *testing.T) {
	tests := []struct {
		path        string
		days        int
		expected    int
		description string
	}{
		{path: "partial_input.txt", days: 80, expected: 5934, description: "Partial input, 80 days"},
		{path: "full_input.txt", days: 80, expected: 345793, description: "Full input, 80 days"},
		//{path: "partial_input.txt", days: 256, expected: 26984457539, description: "Partial input, 256 days"},
		//{path: "full_input.txt", days: 256, expected: 345793, description: "Full input, 256 days"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetLaternfishCount(test.path, test.days)

			if actual != test.expected {
				t.Errorf(
					"GetLaternfishCount(%s, %v) returned %v, expected %v",
					test.path,
					test.days,
					actual,
					test.expected,
				)
			}
		})
	}
}
