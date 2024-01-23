package main

import "sort"
import "fmt"

func main() {
	// fmt.Println(solve([]int{2, 2, 2, 2, 8, 2, 2, 2, 10}))
	fmt.Println(solve([]int{1, 31, 1, 1, 1, 1, 14, 2, 1, 1, 1, 2, 22, 1, 11, 1, 1, 1, 1, 23, 1, 1, 11, 1, 11}))

}

func solve(A []int) []int {
	sort.Ints(A)

	B := make([]int, 0)
	B = append(B, A[len(A)-1])
	ans := B[0]

	for i := len(A) - 2; i >= 0; i-- {
		if !checkPresentInArray(B, A[i]) && A[i] != 1 {
			B = append(B, A[i])
			ans = gcd(ans, A[i])
			fmt.Println(ans)
		}

	}
	return B
}

func checkPresentInArray(ar []int, key int) bool {
	for i := 0; i < len(ar); i++ {
		if ar[i] == key {
			return true
		}
	}
	return false
}

func gcd(A int, B int) int {
	if B == 0 {
		return A
	}
	return gcd(B, A%B)
}
