package main

import "fmt"

func main() {
	var answer int
	answer = solve([]int{4, 3, 2, 6, 1}, 3, 11)
	fmt.Println(answer)
	answer = solve([]int{4, 2, 2, 5, 1}, 4, 6)
	fmt.Println(answer)

}

// sliding window technique
func solve(A []int, B int, C int) int {
	running_sum := 0
	for i := 0; i < B; i++ {
		running_sum += A[i]
	}

	if running_sum == C {
		return 1
	}

	for i := B; i < len(A); i++ {
		running_sum = running_sum + A[i] - A[i-B]
		if running_sum == C {
			return 1
		}
	}

	return 0
}
