package main

import (
	"fmt"
	"math"
)

func main() {
	output := solve([]int{60, 100, 120}, []int{10, 20, 30}, 50)
	fmt.Println(output)
}

func solve(A []int, B []int, C int) int {
	if len(A) == 0 || C == 0 {
		return 0
	} else if B[len(B)-1] <= C {
		a := A[len(A)-1] + solve(A[0:len(A)-1], B[0:len(B)-1], C-B[len(B)-1])
		b := solve(A[0:len(A)-1], B[0:len(B)-1], C)
		out :=  int(math.Max(float64(a), float64(b)))
		return out
	} else if B[len(B)-1] > C {
		solve(A[0:len(A)-1], B[0:len(B)-1], C)
	}
	return 0
}
