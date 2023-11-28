package main

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		description string
		nums1       []int
		nums2       []int
		expected    float64
	}{
		{
			description: "equally long inputs",
			nums1:       []int{1, 2, 4},
			nums2:       []int{1, 3, 4},
			expected:    2.5,
		},
		{
			description: "one input is shorter",
			nums1:       []int{1, 2},
			nums2:       []int{1, 3, 4},
			expected:    2,
		},
		{
			description: "there are negative numbers",
			nums1:       []int{-10 ^ 5, -2, -1},
			nums2:       []int{1, 3, 4},
			expected:    0,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			res := findMedianSortedArrays(test.nums1, test.nums2)

			if res != test.expected {
				t.Errorf("the expected res is %v, received: %v", test.expected, res)
			}
		})
	}
}
