package main

import "fmt"

func main() {
	fmt.Println(solve([]int{6, 4}))
	fmt.Println(solve([]int{2,3,4}))
}

func solve(A []int) int {
	ans := A[0]
	for i := 1; i < len(A); i++ {
		ans = gcd(ans, A[i])
	}
	return ans
}

func gcd(A int, B int) int {
	if B == 0 {
		return A
	}
	return gcd(B, A%B)

}
