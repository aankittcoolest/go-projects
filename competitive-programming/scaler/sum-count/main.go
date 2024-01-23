package main

import "fmt"

func main() {
	fmt.Println("cool")
	output := solve([]int{1, 2, 3, 2, 1}, 5)
	fmt.Println(output)
	output = solve([]int{1, 1, 1}, 2)
	fmt.Println(output)
}

func solve(A []int, B int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i]+A[j] == B {
				count++
			}
		}
	}
	return count

}
