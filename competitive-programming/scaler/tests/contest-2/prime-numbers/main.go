package main

import "fmt"

func main() {
	A := [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(solve(A))

}

func solve(A [][]int) int {
	B := make([]int, 0)
	for i := 0; i < len(A); i++ {
		B = append(B, A[i][i])
		B = append(B, A[i][len(A)-1-i])
	}
	max := -1
	for i := 0; i < len(B); i++ {
		if isPrime(B[i]) && B[i] > max {
			max = B[i]
		}
	}
	return max
}

func isPrime(A int) bool {
	count := 0
	for i := 1; i*i <= A; i++ {
		if A%i == 0 {
			if i == A/i {
				count++

			} else {
				count = count + 2

			}
		}
	}
	if count == 2 {
		return true
	}
	return false

}
