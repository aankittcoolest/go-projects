package main

import "fmt"

var modulo int = 1e9 + 7

func main() {
	var answer int
	answer = solve(2, 27)
	fmt.Println(answer)
	// fmt.Println(answer % modulo)
}

func solve(A int, B int) int {
	return powerMod(A, int(factorial(B))) % modulo
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return (factorial(n-1) * n) % (modulo - 1)
}

func powerMod(a int, n int) int {
	if n == 0 {
		return 1
	}
	y := powerMod(a, n/2)
	if n%2 == 0 {
		return (y * y % modulo) % modulo
	}
	return (y % modulo * y % modulo * a % modulo) % modulo
}
