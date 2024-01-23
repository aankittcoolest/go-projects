package main

import "fmt"

func main() {
	fmt.Println(combine(4, 3))

}

func combine(A int, B int) [][]int {
}

func factorial(A int) int {
	if A == 0 {
		return 1
	}
	return A * factorial(A-1)
}
