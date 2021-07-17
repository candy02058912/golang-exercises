package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name   string
	input  []int
	output []int
}

func TestBubbleSort(t *testing.T) {
	tcs := []testCase{
		{
			name:   "all positive integers",
			input:  []int{9, 2, 5, 4, 3, 7, 6, 8, 1, 10},
			output: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:   "with negative integer",
			input:  []int{9, 2, 5, 4, -3, 7, 6, 8, 1, 10},
			output: []int{-3, 1, 2, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.input, tc.input)
		})
	}
}
