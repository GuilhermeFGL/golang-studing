package addresses

import "testing"

type TestCase struct {
	description string
	input       string
	expected    bool
}

// to run all tests: go test ./... -v
func TestTypeOfAddress(t *testing.T) { // needs to start with Test

	testCases := []TestCase{
		{
			description: "positive case - Street",
			input:       "Street Main number 123",
			expected:    true,
		},
		{
			description: "positive case - Avenue",
			input:       "Avenue Main number 123",
			expected:    true,
		},
		{
			description: "negative case",
			input:       "Highway 66 mile 123",
			expected:    false,
		},
	}

	for _, testCase := range testCases {
		t.Log(testCase.description)
		received := TypeOfAddress(testCase.input)
		if received != testCase.expected {
			t.Errorf("The received type is not the expected type: expected: %t, got: %t", testCase.expected, received)
		}
	}
}
