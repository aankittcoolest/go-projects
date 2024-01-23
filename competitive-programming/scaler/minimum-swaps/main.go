package main

import (
	"fmt"
	"math"
)

func main() {
	var answer int
	// answer = solve([]int{1, 12, 10, 3, 14, 10, 5}, 8)
	// fmt.Println(answer)
	// answer = solve([]int{5, 17, 100, 20, 11}, 20)
	// fmt.Println(answer)

	answer = solve([]int{52, 7, 93, 47, 68, 26, 51, 44, 5, 41, 88, 19, 78, 38, 17, 13, 24, 74, 92, 5, 84, 27, 48, 49, 37, 59, 3, 56, 79, 26, 55, 60, 16, 83, 63, 40, 55, 9, 96, 29, 7, 22, 27, 74, 78, 38, 11, 65, 29, 52, 36, 21, 94, 46, 52, 47, 87, 33, 87, 70}, 19)
	fmt.Println(answer)
}

func solve(A []int, B int) int {
	count := 0
	// find total number of elements less than or equal to B
	for i := 0; i < len(A); i++ {
		if A[i] < B {
			count++
		}
	}

	if count == 0 {
		return 0
	}

	min_bad_counts := math.MaxInt

	bad_counts := 0
	for i := 0; i < count; i++ {
		if A[i] > B {
			bad_counts++

		}
	}
	if min_bad_counts > bad_counts {
		min_bad_counts = bad_counts
	}
	fmt.Println(bad_counts)

	for i := 1; i < len(A)-1; i++ {
		fmt.Println(count+i-1)
		if A[i] > B {
			bad_counts++
		}
		if A[i-1] < B {
			bad_counts--
		}
		if min_bad_counts > bad_counts {
			min_bad_counts = bad_counts
		}
	}
	return min_bad_counts

}
