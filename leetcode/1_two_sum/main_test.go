package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{2, 3, 4, 5},
			target:   9,
			expected: []int{3, 2},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{2, 1},
		},
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 0},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{1, 0},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			res := twoSum(test.nums, test.target)

			if !reflect.DeepEqual(res, test.expected) {
				t.Errorf("the expected res is %q, received: %q", test.expected, res)
			}
		})
	}
}
