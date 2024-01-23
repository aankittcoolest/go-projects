package main

import "fmt"

func main() {
	cntBits([]int{1, 3, 5})
}

func cntBits(A []int) int {
	var one_counts, zero_counts, total_counts int
	mod := 1000000007
	for i := 0; i <= 31; i++ {
		one_counts = 0
		for j := 0; j < len(A); j++ {
			d := A[j] & (1 << i)
			if d > 0 {
				one_counts++
			}
		}
		zero_counts = len(A) - one_counts
		total_counts = total_counts + (2 * one_counts * zero_counts)
		total_counts %= mod

	}
	return total_counts
}
