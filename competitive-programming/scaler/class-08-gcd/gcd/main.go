package main

import "fmt"

func main() {
	fmt.Println(gcd(4,6))
	fmt.Println(gcd(6,7))
}

func gcd(A int, B int) int {
	if B == 0 {
		return A
	}
	return gcd(B, A%B)
}
