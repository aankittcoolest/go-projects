package main

import "fmt"

// problem constraints
var min = 1
var max = 100000 + 1

func main() {
	fmt.Println(solve([]int{3, 2, 5}))
}

func solve(A []int) []int {

	//create array of size max
	B := make([]int, max)

	//create frequency array
	for _, val := range A {
		B[val]++
	}

	// loop through frequency array to re-generate array A
	var k int = 0
	for i, val := range B {
		for j := 1; j <= val; j++ {
			A[k] = i
			k++
		}
	}

	// return sorted array
	return A
}
