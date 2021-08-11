package main

import (
	"fmt"
	"sort"
)

func partialSort(arr []int, ch chan []int) {
	fmt.Println("Sorting subarray:", arr)
	sort.Ints(arr)
	ch <- arr
}

func main() {
	input := []int{3, 4, 1, 2, 3, 2, 9, 8, 7}
	size := len(input) / 4
	ch := make(chan []int)
	for i := 0; i < 4; i++ {
		if i == 3 {
			go partialSort(input[size*(i+1):], ch)
		} else {
			go partialSort(input[size*i:size*(i+1)], ch)
		}
	}
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
