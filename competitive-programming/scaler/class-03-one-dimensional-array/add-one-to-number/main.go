package main

import "fmt"

func main() {
	var answer []int
	answer = plusOne([]int{1, 2, 3})
	fmt.Println(answer)
	answer = plusOne([]int{9, 9, 9})
	fmt.Println(answer)
	answer = plusOne([]int{0, 9, 9})
	fmt.Println(answer)
	answer = plusOne([]int{0, 3, 7, 6, 4, 0, 5, 5, 5})
	fmt.Println(answer)
	answer = plusOne([]int{0, 0, 3, 7, 6, 4, 0, 5, 5, 5})
	fmt.Println(answer)

}

func plusOne(A []int) []int {
	for i := len(A) - 1; i > 0; i-- {
		if A[i] == 9 {
			A[i] = 0
			continue
		}
		A[i] = A[i] + 1
		zero_indexes := -1
		for j := 0; j < len(A); j++ {
			if A[j] == 0 {
				zero_indexes = j
			} else {
				break
			}
		}
		if zero_indexes >= 0 {
			A = A[zero_indexes+1:]
		}
		return A
	}
	if A[0] == 9 {
		A[0] = 0
		A = append([]int{1}, A...)
	} else {
		A[0] = A[0] + 1
	}
	return A
}
