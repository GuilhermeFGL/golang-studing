package addresses

import "testing"

type TestCase struct {
	description string
	input       string
	expected    bool
}

func TestTypeOfAddress(t *testing.T) { // needs to start with Test

	cases := []TestCase{
		{
			description: "positive case",
			input:       "Street Main number 123",
			expected:    true,
		},
		{
			description: "negative case",
			input:       "Highway 66 mile 123",
			expected:    false,
		},
	}

	for _, testCase := range cases {
		t.Log(testCase.description)
		received := TypeOfAddress(testCase.input)
		if received != testCase.expected {
			t.Errorf("The received type is not the expected type: expected: %t, got: %t", testCase.expected, received)
		}
	}
}
