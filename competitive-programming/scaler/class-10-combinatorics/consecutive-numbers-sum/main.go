package main

import "fmt"
import "math"

func main() {
	fmt.Println(solve(5))
	fmt.Println(solve(1))
	fmt.Println(solve(15))
}

func solve(A int) int {
	count := 0
	for i := 1; i <= int(math.Sqrt(2*float64(A))); i++ {
		T := (A - i*(i-1)/2)
		if T%i == 0 {
			count++
		}
	}
	return count
}
