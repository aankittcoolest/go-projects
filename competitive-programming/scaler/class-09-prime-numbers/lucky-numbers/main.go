package main

import "fmt"

func main() {
	fmt.Println(solve(8))
	fmt.Println(solve(12))
}

type luckyFactor struct {
	num           int
	is_pr         bool
	prime_factors []int
}

func solve(A int) int {
	pf := make([]luckyFactor, A+1)

	for i := 0; i < len(pf); i++ {
		pf[i].num = i
		pf[i].is_pr = true
	}

	for i := 2; i < A; i++ {
		if pf[i].is_pr {
			for j := 2 * i; j <= A; j = j + i {
				pf[j].is_pr = false
				pf[j].prime_factors = append(pf[j].prime_factors, i)
			}

		}
	}

	count := 0
	for i := 0; i < len(pf); i++ {
		if len(pf[i].prime_factors) > 1 {
			count++
		}
	}

	return count
}
