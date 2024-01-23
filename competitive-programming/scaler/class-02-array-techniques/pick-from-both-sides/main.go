package main

import (
	"fmt"
	"math"
)

func main() {
	var answer int
	answer = solve([]int{5, -2, 3, 1, 2}, 3)
	fmt.Println(answer)
	answer = solve([]int{2, 3, 4, 3, 4, 4, 1}, 6)
	fmt.Println(answer)
}

func solve(A []int, B int) int {

	C := make([]int, 0)
	for i := B - 1; i >= 0; i-- {
		C = append(C, A[i])
	}

	for i := len(A) - 1; i >= len(A)-B; i-- {
		C = append(C, A[i])
	}

	sum := 0
	max_sum := math.MinInt
	//calculate max sum using sliding window
	for i := 0; i < len(C); i++ {
		if i > B-1 {
			if max_sum < sum {
				max_sum = sum
			}
			sum -= C[i-B]
		}
		sum += C[i]
	}

	return max_sum
}
