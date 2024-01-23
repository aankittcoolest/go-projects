package main

import "fmt"

func main() {
	fmt.Println(findRank("acb"))
	fmt.Println(findRank("a"))
	fmt.Println(findRank("VIEW"))
	// fmt.Println(factorial(5))

}

var arr = make([]int, 1001)
var mod = 1000003

func findRank(A string) int {
	B := []rune(A)

	ans := 0
	for i := 0; i < len(B)-1; i++ {
		count := 0
		for j := i + 1; j < len(B); j++ {
			if B[j] < B[i] {
				count++
			}
		}

		if count > 0 {
			ans = ans + factorial(len(B)-i-1)*count
		}
	}

	return ans + 1
}

func factorial(A int) int {
	if A == 0 || A == 1 {
		return 1
	}
	if arr[A] != 0 {
		return arr[A]
	}
	arr[A] = A * factorial(A-1) % mod
	return arr[A]
}
