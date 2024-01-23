package main

import "fmt"

func main() {
	var answer int
	answer = maxSubArray([]int{1, 2, 3, 4, -10})
	fmt.Println(answer)
	answer = maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(answer)
	answer = maxSubArray([]int{-163, -20})
	fmt.Println(answer)

}

// kadane's algorithm
func maxSubArray(A []int) int {
	answer := A[0]
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		if sum > answer {
			answer = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return answer

}
