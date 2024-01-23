package main

import "fmt"

func main() {
	fmt.Println(solve([]int{1, 2, 3}))
	fmt.Println(solve([]int{28, 7, 3, 6, 23, 16, 5, 29, 23}))

}

func solve(A []int) int {
	ans := 0
	for i := 0; i < 32; i++ {
		set := 0
		unset := 0
		for j := 0; j < len(A); j++ {
			if ((A[j] >> i) & 1) == 1 {
				set++
			} else {
				unset++
			}
		}
		ans = (ans + (set * unset * (1 << i))) % 1000000007
	}

	return ans
}
