package main

import "fmt"

// https://www.scaler.com/academy/mentee-dashboard/class/81671/homework/problems/1085/

/*
   Three observations
   1. Every even number can be splitted into two equal numbers and hence their XOR is always zero.
   2. Every odd number can be splitted into a combination of even and odd number.
   3. Sum of two odd numbers is always even, hence their XOR is always zero.
*/

func main() {
	var ans string
	ans = solve([]int{9, 17})
	fmt.Println(ans)
	ans = solve([]int{1})
	fmt.Println(ans)
}

func solve(A []int) string {
	odd_count := 0
	for _, element := range A {
		if element%2 != 0 {
			odd_count++
		}
	}
	if odd_count%2 == 0 {
		return "Yes"
	}
	return "No"
}
