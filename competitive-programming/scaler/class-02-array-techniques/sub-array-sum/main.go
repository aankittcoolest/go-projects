package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 2, 3}))
}

func subarraySum(A []int) int64 {
	var total_sum int64
	// use contributon forumula = (i+1) * (n-i)
	for i := 0; i < len(A); i++ {
		contribution := (i + 1) * (len(A) - i)
		total_sum += int64(A[i] * contribution)
	}
	return total_sum
}
