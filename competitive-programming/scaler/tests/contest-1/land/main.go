package main

import "fmt"

func main() {

	A := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	// A := [][]int{
	// 	{1},
	// }
	result := solve(A)
	fmt.Println(result)
}

func solve(A [][]int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == 1 {
				if i == 0 {
					count++
				}
				if i == len(A)-1 {
					count++

				}
				if j == 0 {
					count++
				}
				if j == len(A[i])-1 {
					count++
				}
				if i > 0 {
					if A[i-1][j] == 0 {
						count++
					}
				}
				if i < len(A)-1 {
					if A[i+1][j] == 0 {
						count++
					}
				}
				if j > 0 {
					if A[i][j-1] == 0 {
						count++
					}
				}
				if j < len(A[i])-1 {
					if A[i][j+1] == 0 {
						count++
					}
				}
			}
		}
	}
	return count
}
