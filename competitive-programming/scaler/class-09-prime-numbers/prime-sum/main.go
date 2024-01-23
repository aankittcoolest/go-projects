package main

import "fmt"

func main() {
	fmt.Println(primesum(10))
}

func primesum(A int) []int {
	pf := make([]bool, A+1)
	for i := 0; i < len(pf); i++ {
		pf[i] = true
	}

	for i := 2; i*i < A; i++ {
		for j := i * i; j <= A; j = j + i {
			pf[j] = false
		}
	}

	C := make([]int, 0)
	for i := 2; i < len(pf); i++ {
		if pf[i] == true {
			C = append(C, i)
		}
	}

	i := 0
	j := len(C) - 1

	for {
		if C[i]+C[j] == A {
			return []int{C[i], C[j]}
		}
		if C[i]+C[j] > A {
			j--
		}
		if C[i]+C[j] < A {
			i++
		}
	}

}
