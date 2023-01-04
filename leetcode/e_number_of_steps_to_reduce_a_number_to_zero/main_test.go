package main

import (
	"fmt"
	"testing"
)

func TestNumberOfSteps(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{in: 14, out: 6},
		{in: 8, out: 4},
		{in: 123, out: 12},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.in), func(t *testing.T) {
			got := numberOfSteps(test.in)
			if got != test.out {
				t.Errorf("numberOfSteps(%v) = %v, want %v", test.in, got, test.out)
			}
		})
	}
}
