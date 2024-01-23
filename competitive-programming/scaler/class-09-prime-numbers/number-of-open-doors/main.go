package main

import (
	"fmt"
	"math"
)

func main() {
	// solve(6)
	// solve(5)
	numberOfPerfectSquares(5)
	numberOfPerfectSquares(6)
	numberOfPerfectSquares(100)
}

func solve(A int) int {
	pf := make([]bool, A+1)
	for i := 0; i < len(pf); i++ {
		pf[i] = true
	}

	for i := 2; i <= A; i++ {
		pf[i] = !pf[i]
		for j := 2 * i; j <= A; j = j + i {
			pf[j] = !pf[j]
		}
	}
	count := 0
	for i := 1; i < len(pf); i++ {
		if pf[i] == true {
			count++
		}
	}
	return count
}

func numberOfPerfectSquares(A int) int {
	left := 0
	right := A

	var mid int
	for {

		mid = int(math.Ceil(float64(left+right) / 2))
		if mid*mid > A {
			right = mid - 1
		} else if mid*mid < A {
			left = mid + 1
		}
		if left > right || mid*mid == A {
			break
		}
	}
	fmt.Println(mid)
	return mid
}
