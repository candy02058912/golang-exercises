package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name   string
	input  []float64
	output float64
}

func TestDisplaceFn(t *testing.T) {
	tcs := []testCase{
		{
			name:   "acceration 10, initialV 2, initialD 0, time 3",
			input:  []float64{10, 2 , 0, 3},
			output: 51,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
      displaceFn := GenDisplaceFn(tc.input[0], tc.input[1], tc.input[2])
			assert.Equal(t, displaceFn(tc.input[3]), tc.output)
		})
	}
}
