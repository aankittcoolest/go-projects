package main

import (
	"fmt"
)

func main() {
	fmt.Println("cool")
	output := solve([]int{0, 2, 9})
	fmt.Println(output)
	// output = solve([]int{1, 1, 1}, 2)
	// fmt.Println(output)
}

func solve(A []int )  (int) {
	max_even := -1000000000
	min_odd := 1000000000
	for i:=0 ;i < len(A);i ++ {
		if A[i] % 2 == 0  && A[i] > max_even{
			max_even = A[i]
		} else if A[i] < min_odd {
			min_odd = A[i]
		}
	}
	return max_even - min_odd
}