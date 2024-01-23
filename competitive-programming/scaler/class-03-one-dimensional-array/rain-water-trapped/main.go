package main

import "math"
import "fmt"

func main() {
	var answer int
	answer = trap([]int{0, 1, 0, 2})
	fmt.Println(answer)
	// 	trap([]int{4, 2, 5, 7, 4, 2, 3, 6, 8, 2, 3})

}

func trap(A []int) int {
	lmax := make([]int, len(A))
	rmax := make([]int, len(A))

	lmax[0] = A[0]
	for i := 1; i < len(A); i++ {
		lmax[i] = int(math.Max(float64(lmax[i-1]), float64(A[i])))
	}
	rmax[len(A)-1] = A[len(A)-1]
	for i := len(A) - 2; i >= 0; i-- {
		rmax[i] = int(math.Max(float64(rmax[i+1]), float64(A[i])))
	}
	water := 0
	for i := 0; i < len(A); i++ {
		min := int(math.Min(float64(lmax[i]), float64(rmax[i])))
		water = water + (min - A[i])
	}
	return water
}
