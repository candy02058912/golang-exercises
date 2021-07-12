package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter 10 integers, space seperated, e.g. 3 9 7 9 -1 2 1 0 3 -6")
	input := make([]int, 10)
	for i := range input {
		_, err := fmt.Scan(&input[i])
		if err != nil {
			fmt.Errorf("error with input %v: %v", input[:i], err)
		}
	}
  fmt.Println("Received input:", input)
	BubbleSort(input)
	fmt.Println("Sorted output:", input)
}

func BubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func Swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]
}
