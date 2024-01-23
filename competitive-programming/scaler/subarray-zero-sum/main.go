package main

import "fmt"

func main() {
	output := solve([]int{-1, 1})
	fmt.Println(output)
	output = solve([]int{1,2,3,4,5})
	fmt.Println(output)
}

func solve(A []int) int {
	//prefix sum
	sum := 0
	prefix := make(map[int]int)

	for i:=0; i< len(A);i++ {
		sum += A[i]
		if sum == 0 {
			return 1
		} 

		_,ok := prefix[sum]; if ok {
			return 1
		}
		prefix[sum] = 1
	}
	return 0
	/*
	   Just write your code below to complete the function. Required input is available to you as the function arguments.
	   Do not print the result or any output. Just return the result via this function.
	*/

}