package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{
			ransomNote: "a",
			magazine:   "b",
			expected:   false,
		},
		{
			ransomNote: "a",
			magazine:   "a",
			expected:   true,
		},
		{
			ransomNote: "well done",
			magazine:   "hello world",
			expected:   false,
		},
		{
			ransomNote: "well hood",
			magazine:   "hello world",
			expected:   true,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := canConstruct(test.ransomNote, test.magazine)

			if res != test.expected {
				t.Errorf("the expected res is %v, received: %v", test.expected, res)
			}
		})
	}
}
