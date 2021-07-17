package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	fn := GenDisplaceFn(3, 4, 5)
	fmt.Println(fn(5))
}

func GenDisplaceFn(acceleration, velocity, displacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return 3
	}
}
