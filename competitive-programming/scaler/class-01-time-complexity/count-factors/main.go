package main

import (
	"fmt"
	"math"
)

func main() {
	var ans int
	ans = solve(360)
	fmt.Println(ans)
}

func solve(A int) int {
	count := 0

	B := int(math.Sqrt(float64(A)))
	for i := 1; i <= B; i++ {
		if i*i == A {
			count++
		} else if A%i == 0 {
			count = count + 2
		}
	}
	return count
}
