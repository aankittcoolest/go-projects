package main

import "fmt"

func main() {
	arrange([]int{1, 0})
	arrange([]int{0, 2, 1, 3})
	arrange([]int{1, 2, 7, 0, 9, 3, 6, 8, 5, 4})

}

func arrange(A []int) {
	for i := 0; i < len(A); i++ {
		A[i] = A[i] + (A[A[i]]%len(A))*len(A)
	}
	for i := 0; i < len(A); i++ {
		A[i] = A[i] / len(A)
	}
	for i := 0; i < len(A); i++ {
		fmt.Printf("%v ", A[i])
	}
	fmt.Println()
}
