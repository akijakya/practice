package main

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		roman     string
		expected int
	}{
		{
			roman: "III",
			expected: 3,
		},
		{
			roman: "LVIII",
			expected: 58,
		},
		{
			roman: "MCMXCIV",
			expected: 1994,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := romanToInt(test.roman)

			if res != test.expected {
				t.Errorf("the expected res is %q, received: %q", test.expected, res)
			}
		})
	}

}
