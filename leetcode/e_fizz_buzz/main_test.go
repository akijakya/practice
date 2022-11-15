package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n        int
		expected []string
	}{
		{
			n:        5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			n:        15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := fizzBuzz(test.n)

			if !reflect.DeepEqual(res, test.expected) {
				t.Errorf("the expected res is %v, received: %v", test.expected, res)
			}
		})
	}
}
