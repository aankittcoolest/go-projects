package main

import "fmt"

func main() {
	fmt.Println(solve([]int{2, 3, 4, 5}))
	fmt.Println(solve([]int{8, 9, 10}))
}

func solve(A []int) []int {
	pf := createPrimeFactorsArray(findMax(A))
	for i := 0; i < len(A); i++ {
		m := make(map[int]int)
		for {
			if A[i] == 1 {
				break
			}
			factor := pf[A[i]]
			_, ok := m[factor]
			if ok {
				m[factor] = m[factor] + 1
			} else {
				m[factor] = 1

			}
			A[i] = A[i] / factor
		}

		//count number of divisors from map
		ans := 1
		for _, v := range m {
			ans += ans * v
		}
		A[i] = ans
	}
	return A
}

func findMax(A []int) int {
	max := A[0]
	for i := 1; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
		}
	}
	return max
}

func createPrimeFactorsArray(max int) []int {
	pf := make([]int, max+1)

	for i := 0; i < len(pf); i++ {
		pf[i] = i
	}

	for i := 2; i*i < len(pf); i++ {
		if pf[i] == i {
			for j := i * i; j < len(pf); j = j + i {
				pf[j] = i
			}
		}
	}
	return pf
}
