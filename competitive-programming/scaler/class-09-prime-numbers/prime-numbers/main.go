package main

import "fmt"

func main() {
	fmt.Println(solve(7))
	fmt.Println(solve(12))

}

func solve(A int) []int {
	B := make([]bool, A+1)

	for i := 0; i < len(B); i++ {
		B[i] = true
	}

	for i := 2; i*i < len(B); i++ {
		if B[i] == true {
			for j := i * i; j < len(B); j = j + i {
				B[j] = false
			}
		}
	}
	C := make([]int, 0)
	for i := 2; i < len(B); i++ {
		if B[i] == true {
			C = append(C, i)
		}
	}
	return C
}
