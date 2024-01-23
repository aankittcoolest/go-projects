package main

import "fmt"

func main() {
	var answer int
	answer = solve([][]int{
		{1, 1},
		{1, 1},
	})
	fmt.Println(answer)
}

// we will use contribution technique to solve this problem
// number of times an element contributes in a 2D array is repsented by:
// contribution = (i+1)*(j+1) * (n-i)*(j-i)
// hence total sum contributed by an element
// element * contribution
func solve(A [][]int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			contribution := (i + 1) * (j + 1) * (len(A) - i) * (len(A[i]) - j)
			sum += A[i][j] * contribution
		}
	}
	return sum
}
