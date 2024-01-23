package main

import (
	"fmt"
)

func main() {
	output := solve([]int{3,34,4,12,5,2}, 9)
	fmt.Println(output)
	output = solve([]int{3, 34, 4, 12, 5, 2}, 30)
	fmt.Println(output)
}

func solve(A []int, B int) int {
	output := subset_sum(A, B)
	if output {
		return 1
	}
	return 0
}

func subset_sum(A []int, B int) bool {
	if len(A) == 0 && B != 0 {
		return false
	} else if len(A) == 0 || B == 0 {
		return true
	} 
	if A[len(A)-1] <= B {
		return subset_sum(A[:len(A)-1], B-A[len(A)-1]) || subset_sum(A[:len(A)-1], B)
	}
	return subset_sum(A[:len(A)-1], B)
}
