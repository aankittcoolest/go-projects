package main

import "fmt"

func main() {
	var answer int
	answer = solve([]int{53, 39, 88})
	fmt.Println(answer)
	answer = solve([]int{38, 44, 84, 12})
	fmt.Println(answer)
	answer = solve([]int{98, 12, 44, 24, 1, 34, 70, 74, 67, 18, 54, 14})
	fmt.Println(answer)
}

func solve(A []int) int {

	B := make([]int, 0)
	for i := 31; i >= 0; i-- {
		B = nil
		for j := 0; j < len(A); j++ {
			if ((A[j] >> i) & 1) == 1 {
				B = append(B, A[j])
			}
		}
		if len(B) == 2 {
			return B[0] & B[1]
		}
		if len(B) == 1 {
			continue
		}
		if len(B) > 2 {
			A = B
		}
	}
	return A[0] & A[1]
}
