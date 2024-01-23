package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 2, 3, 1}))

}

func singleNumber(A []int) int {
	ans := A[0]
	for i := 1; i < len(A); i++ {
		ans = ans ^ A[i]
	}
	return ans
}
