package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter values for acceleration, initial velocity, and initial displacement, space seperated, e.g. 10 2 1")
	input := ReadInput(3)
	acc, initialV, initialD := input[0], input[1], input[2]
	displaceFn := GenDisplaceFn(acc, initialV, initialD)
	fmt.Println("Enter a value for time, e.g. 3")
	input = ReadInput(1)
	time := input[0]
	fmt.Printf("Displacement: %v\n", displaceFn(time))
}

func ReadInput(size int) []float64 {
	input := make([]float64, size)
	for i := range input {
		_, err := fmt.Scan(&input[i])
		if err != nil {
			fmt.Errorf("error with input %v: %v", input[:i], err)
		}
	}
	return input
}

func GenDisplaceFn(acc, initialV, initialD float64) func(float64) float64 {
	return func(time float64) float64 {
		return 1/2.0*acc*math.Pow(time, 2) + initialV*time + initialD
	}
}
