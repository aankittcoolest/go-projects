package main

import "fmt"

func main() {
	var answer []int
	answer = solve([]int{1, 2, 3, 4, 5})
	fmt.Println(answer)
}

func solve(A []int) []int {
	product := A[0]
	for i := 1; i < len(A); i++ {
		product *= A[i]
	}
	for i := 0; i < len(A); i++ {
		A[i] = product / A[i]
	}
	return A
}
