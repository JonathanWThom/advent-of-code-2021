package day_7

import "testing"

func TestGetFuel(t *testing.T) {
	tests := []struct {
		path        string
		expected    int
		description string
	}{
		{path: "partial_input.txt", expected: 37, description: "Partial input"},
		{path: "full_input.txt", expected: 349769, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetFuel(test.path)

			if actual != test.expected {
				t.Errorf(
					"GetFuel(%s) returned %v, expected %v",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}

func TestGetNonConstantFuel(t *testing.T) {
	tests := []struct {
		path        string
		expected    int
		description string
	}{
		{path: "partial_input.txt", expected: 168, description: "Partial input"},
		{path: "full_input.txt", expected: 99540554, description: "Full input"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			actual := GetNonConstantFuel(test.path)

			if actual != test.expected {
				t.Errorf(
					"GetNonConstantFuel(%s) returned %v, expected %v",
					test.path,
					actual,
					test.expected,
				)
			}
		})
	}
}
