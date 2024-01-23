package main

import "fmt"

func main() {
	// fmt.Println(gcd(2, 8))
	fmt.Println(gcd(8, 10))

}

func solve(A int, B int, C int) int {
	return gcd(A, B*C)
}

func gcd(A int, B int) int {
	if B == 0 {
		fmt.Println(A)
		return A
	}
	c := gcd(B, A%B)
	fmt.Println(c)
	return c
}
