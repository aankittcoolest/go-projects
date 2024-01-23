package main

import (
	"fmt"
	"math"
)

func main() {
	miniMaxSum([]int32{1, 3, 5, 7, 9})
}

func miniMaxSum(arr []int32) {
	var min int32 = math.MaxInt32
	var max int32 = 0
	var sum int64 = 0
	for _, val := range arr {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
		sum += int64(val)
	}
	fmt.Println(sum-int64(max), sum-int64(min))
}
