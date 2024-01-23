package main

import "fmt"

func main() {
	fmt.Println(solve([]int{4, 7, 9}, []int{2, 11, 19}))
	fmt.Println(solve([]int{1}, []int{2}))
}

// two pointer approach
func solve(A []int, B []int) []int {
	C := make([]int, len(A)+len(B))
	P1 := 0
	P2 := 0
	k := 0
	for {
		if P1 >= len(A) || P2 >= len(B) {
			break
		}
		if A[P1] < B[P2] {
			C[k] = A[P1]
			P1++
		} else {
			C[k] = B[P2]
			P2++
		}
		k++
	}

	for {
		if P1 >= len(A) {
			break
		}
		C[k] = A[P1]
		P1++
		k++
	}
	for {
		if P2 >= len(B) {
			break
		}
		C[k] = B[P2]
		P2++
		k++
	}
	return C
}
