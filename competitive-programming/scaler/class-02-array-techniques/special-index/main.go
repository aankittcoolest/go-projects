package main

import "fmt"

func main() {
	var answer int
	answer = solve([]int{2, 1, 6, 4})
	fmt.Println(answer)
	answer = solve([]int{1, 1, 1})
	fmt.Println(answer)
}

func solve(A []int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		even_sum := 0
		odd_sum := 0

		for j := 0; j < len(A); j++ {
			if i == j {
				continue
			}
			k := 0
			if j < i {
				k = j
			} else {
				k = j - 1

			}
			// even number
			if k%2 == 0 {
				even_sum = even_sum + A[j]
			} else {
				odd_sum = odd_sum + A[j]
			}
		}
		if even_sum == odd_sum {
			count++
		}
	}
	return count
}
