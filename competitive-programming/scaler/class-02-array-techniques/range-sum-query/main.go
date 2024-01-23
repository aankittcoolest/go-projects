package main

import "fmt"

func main() {
	var output []int64
	output = rangeSum([]int{1, 2, 3, 4, 5}, [][]int{{0, 3}, {1, 2}})
	fmt.Println(output)
}

func rangeSum(A []int, B [][]int) []int64 {
	prefix_sum_array := make([]int, len(A))
	output := make([]int64, len(B))
	prefix_sum_array[0] = A[0]
	for i := 1; i < len(A); i++ {
		prefix_sum_array[i] = prefix_sum_array[i-1] + A[i]
	}

	for i := 0; i < len(B); i++ {
		if B[i][0] == 0 {
			output[i] = int64(prefix_sum_array[B[i][1]])
		} else {
			output[i] = int64(prefix_sum_array[B[i][1]] - prefix_sum_array[B[i][0]-1])
		}

	}
	return output
}
