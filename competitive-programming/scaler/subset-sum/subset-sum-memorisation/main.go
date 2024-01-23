package main

import "fmt"

var item_range = 100
var sum_range = 100000
var t [102][100002]track

type track struct {
	item      bool
	traversed bool
}

func main() {
	for i := 0; i < item_range+2; i++ {
		for j := 0; j < sum_range+2; j++ {
			t[i][j] = track{
				false,
				false,
			}
		}
	}

	output := solve([]int{3, 34, 4, 12, 5, 2}, 9)
	fmt.Println(output)
	output = solve([]int{3, 34, 4, 12, 5, 2}, 30)
	fmt.Println(output)
}

func solve(A []int, B int) int {
	output := subset_sum(A, B)
	if output {
		return 1
	}
	return 0
}

func subset_sum(A []int, B int) bool {
	if len(A) == 0 && B != 0 {
		return false
	} else if len(A) == 0 || B == 0 {
		return true
	}
	if t[len(A)][B].traversed {
		return t[len(A)][B].item
	}
	if A[len(A)-1] <= B {
		t[len(A)][B].item = subset_sum(A[:len(A)-1], B-A[len(A)-1]) || subset_sum(A[:len(A)-1], B)
		t[len(A)][B].traversed = true
		return t[len(A)][B].item
	}
	t[len(A)][B].item = subset_sum(A[:len(A)-1], B)
	t[len(A)][B].traversed = true
	return t[len(A)][B].item
}
