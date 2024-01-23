package main

import "math"
import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 4, 3, 3, 2, 2, 3, 1, 1}))
	fmt.Println(singleNumber([]int{0, 0, 0, 1}))
}

func singleNumber(A []int) int {
	ans := 0
	for i := 0; i < 31; i++ {
		set_count := 0
		for j := 0; j < len(A); j++ {
			if A[j]>>i&1 == 1 {
				set_count++
			}
		}
		ans = ans + set_count%3*int(math.Pow(float64(2), float64(i)))
	}
	return ans
}
