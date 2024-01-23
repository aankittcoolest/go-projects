package main

import "fmt"

func main() {
	fmt.Println(solve([]int{1, 2, 3}, []int{0, 1}))
	fmt.Println(solve([]int{2, 4, 7, 11}, []int{3}))

}

func solve(A []int, B []int) []int {
	for i := 0; i < len(B); i++ {
		set_count := 0
		unset_count := 0
		for j := 0; j < len(A); j++ {
			if (A[j]>>B[i])&1 == 1 {
				set_count++
			} else {
				unset_count++
			}
		}
		B[i] = set_count * unset_count
	}
	return B
}
