package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partialSort(arr []int, ch chan []int) {
	fmt.Println("Sorting subarray:", arr)
	sort.Ints(arr)
	ch <- arr
}

func mergeTwoSlices(s1, s2 []int) []int {
	p1, p2 := 0, 0
	output := []int{}
	for p1 < len(s1) && p2 < len(s2) {
		if s1[p1] <= s2[p2] {
			output = append(output, s1[p1])
			p1 += 1
		} else {
			output = append(output, s2[p2])
			p2 += 1
		}
	}
	if p1 < len(s1) {
		output = append(output, s1[p1:]...)
	}
	if p2 < len(s2) {
		output = append(output, s2[p2:]...)
	}
	return output
}

func readInput() ([]int, error) {
	fmt.Println("Enter integers to be sorted, space seperated, e.g. 3 9 7 9 -1 2 1 0 3 -6 -8 -10 3")
	var input []int
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		rawInput := strings.Split(scanner.Text(), " ")
		for _, n := range rawInput {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, fmt.Errorf("error parsing input: %v", err)
			}
			input = append(input, num)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Errorf("error reading input: %v", err)
	}
	return input, nil
}

func main() {
	input, err := readInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
  // partition size for the 4 subarrays
	size := len(input) / 4
	ch := make(chan []int)

	for i := 0; i < 4; i++ {
		if i == 3 {
			// put the rest of the numbers inside the last partition
			go partialSort(input[size*i:], ch)
		} else {
			go partialSort(input[size*i:size*(i+1)], ch)
		}
	}

	p1 := mergeTwoSlices(<-ch, <-ch)
	p2 := mergeTwoSlices(<-ch, <-ch)
	output := mergeTwoSlices(p1, p2)

	fmt.Println("Sorted array:", output)
}
