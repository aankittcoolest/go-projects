package main

import "fmt"

func main() {
	fmt.Println(solve([]int{1, 2, 3}))

}

func solve(A []int) int {
	sum := 0
	freq := make([]int, 10001)
	freq[0] = 0
	for i := 0; i < len(A); i++ {
		freq[A[i]] = freq[A[i]] + 1
	}
	fmt.Println(freq)
	for i := 1; i < len(freq); i++ {
		for j := 1; j < len(freq); j++ {
			mod := i % j
			temp := freq[i] * freq[j] * mod

			sum = (sum%1000000007 + temp%1000000007)
		}
	}
	return sum
}
