package main

import "math"
import "fmt"

func main() {
	var answer int
	answer = solve([]int{3, 7, 90, 20, 10, 50, 40}, 3)
	fmt.Println(answer)
	answer = solve([]int{3, 7, 5, 20, -10, 0, 12}, 2)
	fmt.Println(answer)
	answer = solve([]int{18, 11, 16, 19, 11, 9, 8, 15, 3, 10, 9, 20, 1, 19}, 1)
	fmt.Println(answer)
	answer = solve([]int{5, 15, 7, 6, 3, 4, 18, 14, 13, 7, 3, 7, 2, 18}, 6)
	fmt.Println(answer)
	answer = solve([]int{20, 3, 13, 5, 10, 14, 8, 5, 11, 9, 1, 11}, 9)
	fmt.Println(answer)
	answer = solve([]int{3, 16, 11, 13, 19, 17, 11, 4, 9, 9, 6, 7, 3, 6, 12, 3, 4, 15, 5, 19}, 1)
	fmt.Println(answer)
}

func solve(A []int, B int) int {
	index := 0
	min := math.MaxInt
	sum := 0
	for i := 0; i < len(A); i++ {
		if i >= B {
			if sum < min {
				index = i - B
				min = sum
			}
			sum = sum - A[i-B]
		}
		sum += A[i]
	}
	if sum < min {
		index = len(A) - B
	}
	return index
}
