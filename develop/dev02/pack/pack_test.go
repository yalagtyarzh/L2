package pack

import "testing"

func Test_Unpack(t *testing.T) {
	theTests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "valid",
			input:    "a4bc2d5e\\\\",
			expected: "aaaabccddddde\\",
		},
		{
			name:     "invalid unshielded number",
			input:    "45",
			expected: "",
		},
		{
			name:     "invalid backslash in end of string",
			input:    "a4bc2d5e\\",
			expected: "",
		},
	}

	for _, tt := range theTests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, _ := Unpack(tt.input)

				if got != tt.expected {
					t.Errorf("Expected: %s. Got: %s", tt.expected, got)
				}
			},
		)
	}
}
