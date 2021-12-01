package day_1

import "testing"

func TestGetDepthIncreaseCount(t *testing.T) {
	tests := []struct {
		path        string
		expected    int
		description string
	}{
		{path: "partial_input.txt", expected: 7, description: "Partial input"},
		{path: "full_input.txt", expected: 1466, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetDepthIncreaseCount(test.path)

			if actual != test.expected {
				t.Errorf(
					"GetDepthIncreaseCount(%s) returned %v, expected %v",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}

func TestGetSlidingDepthIncreaseCount(t *testing.T) {
	tests := []struct {
		path        string
		expected    int
		description string
	}{
		{path: "partial_input.txt", expected: 5, description: "Partial input"},
		{path: "full_input.txt", expected: 1491, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetSlidingDepthIncreaseCount(test.path)

			if actual != test.expected {
				t.Errorf(
					"GetSlidingDepthIncreaseCount(%s) returned %v, expected %v",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}
