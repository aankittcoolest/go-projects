package main

import "fmt"
import "math"

func main() {
	fmt.Println(solve(1, 2))
	fmt.Println(solve(5, 10))
}

func solve(A int, B int) int {
	return int(math.Abs(float64(A - B)))
}
