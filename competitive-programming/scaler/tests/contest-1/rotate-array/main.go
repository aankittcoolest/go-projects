package main

import "fmt"

func main() {
	var answer int
	answer = solve([]int{3, 4, 1, 2})
	fmt.Println(answer)
	answer = solve([]int{2, 1, 9, 9})
	fmt.Println(answer)
	answer = solve([]int{7, 10})
	fmt.Println(answer)
	answer = solve([]int{14, 189, 762, 1, 172, 673})
	fmt.Println(answer)
}

func solve(A []int) int {
	index := -1
	for i := 0; i < len(A)-1; i++ {
		if A[i+1] >= A[i] {
			continue
		} else {
			index = i
			break
		}
	}

	if index == -1 {
		return 1
	}

	for i := index + 1; i < len(A)-1; i++ {
		if A[i] <= A[i+1] && A[i] <= A[0] && A[i+1] <= A[0] {
			continue
		} else {
			return 0
		}
	}

	return 1
}
