package main

import "fmt"

func main() {
	// solve([][]int{
	// 	{3, 2, 4, 1},
	// 	{-1, 4, 3, 2},
	// 	{2, 7, 6, 3},
	// })

	var answer int64
	answer = solve([][]int{
		{-5, -4, -3},
		{-1, 2, 3},
		{2, 2, 4},
	})
	fmt.Println(answer)

	//	answer = solve([][]int{
	//		{1, 2, 3},
	//		{4, 5, 6},
	//		{7, 8, 9},
	//	})
	//
	// fmt.Println(answer)
}

func calculate_psum(A [][]int) [][]int64 {
	B := make([][]int64, len(A))

	for i := range B {
		B[i] = make([]int64, len(A[0]))
	}

	for i := range A {
		B[i][0] = int64(A[i][0])
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if j > 0 {
				B[i][j] = int64(A[i][j]) + B[i][j-1]
			}
		}
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if i > 0 {
				B[i][j] = int64(B[i][j]) + B[i-1][j]
			}
		}
	}
	return B
}

func solve(A [][]int) int64 {
	B := calculate_psum(A)
	fmt.Println(B)

	max_answer := B[len(B)-1][len(B[0])-1]
	for i := 0; i < len(B); i++ {
		for j := 0; j < len(B[i]); j++ {
			ans := B[len(B)-1][len(B[0])-1]
			if j > 0 {
				ans = ans - B[len(B)-1][j-1]
			}
			if i > 0 {
				ans = ans - B[i-1][len(B[0])-1]
			}
			if i > 0 && j > 0 {
				ans = ans + B[i-1][j-1]
			}
			if max_answer < ans {
				max_answer = ans
			}
		}
	}
	return max_answer
}
