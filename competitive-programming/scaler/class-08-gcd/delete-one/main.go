package main

import "fmt"

func main() {
	fmt.Println(solve([]int{12, 15, 18}))

}

func solve(A []int) int {
	prefix_array := make([]int, len(A))
	suffix_array := make([]int, len(A))

	prefix_array[0] = A[0]
	suffix_array[len(A)-1] = A[len(A)-1]

	for i := 1; i < len(A); i++ {
		prefix_array[i] = gcd(prefix_array[i-1], A[i])
	}

	for i := len(A) - 2; i > 0; i-- {
		suffix_array[i] = gcd(suffix_array[i+1], A[i])
	}

	max_gcd := 0
	current_gcd := 0
	for i := 0; i < len(A); i++ {
		if i == 0 {
			current_gcd = suffix_array[1]
		} else if i == len(A)-1 {
			current_gcd = prefix_array[len(A)-2]
		} else {
			current_gcd = gcd(prefix_array[i-1], suffix_array[i+1])
		}
		if current_gcd > max_gcd {
			max_gcd = current_gcd
		}
	}

	return max_gcd
}

func gcd(A int, B int) int {
	if B == 0 {
		return A
	}
	return gcd(B, A%B)
}
