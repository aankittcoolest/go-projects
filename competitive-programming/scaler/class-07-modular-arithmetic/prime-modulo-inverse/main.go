package main

import "fmt"

func main() {
	fmt.Println(solve(3, 5))
	fmt.Println(solve(6, 23))
	fmt.Println(solve(21, 17))

}

// we will use fermat theorem
// A ^ (-1) mod B = A ^ (B-2) % B
func solve(A int, B int) int {
	return powerMod(A, B-2, B)
}

func powerMod(a int, n int, modulo int) int {
	if n == 0 {
		return 1
	}
	y := powerMod(a, n/2, modulo)
	if n%2 == 0 {
		return (y * y % modulo) % modulo
	}
	return (y % modulo * y % modulo * a % modulo) % modulo
}
