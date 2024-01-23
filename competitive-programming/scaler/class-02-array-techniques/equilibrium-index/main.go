package main

import (
	"fmt"
	"math"
)

func main() {
	solve([]int{-7, 1, 5, 2, -4, 3, 0})
}

// solve ...
func solve(A []int) int {
	prefix_sum_array := make([]int64, len(A))
	prefix_sum_array[0] = int64(A[0])
	for i := 1; i < len(A); i++ {
		prefix_sum_array[i] = prefix_sum_array[i-1] + int64(A[i])
	}

	for i := 0; i < len(A); i++ {
		// lsum = sum(0, i-1)
		// rsum = sum(i+1, N-1)
		var lsum, rsum int64
		lsum = math.MinInt64
		rsum = math.MinInt64
		if i == 0 {
			lsum = 0
			rsum = prefix_sum_array[len(A)-1] - prefix_sum_array[i]
		} else if i == len(A)-1 {
			lsum = prefix_sum_array[i-1]
			rsum = 0
		} else {
			lsum = prefix_sum_array[i-1]
			rsum = prefix_sum_array[len(A)-1] - prefix_sum_array[i]
		}
		if lsum == rsum {
			return i
		}

	}
	return -1
}
