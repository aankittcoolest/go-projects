package main

import "math"

func main() {
	diagonalDifference([][]int32{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	//fmt.Println([][]int32{{1, 2, 3}, {4, 5, 6}, {7,8,9}})

}
func diagonalDifference(arr [][]int32) int32 {
	t := len(arr)
	var sum1, sum2 int32 = 0, 0

	for i := 0; i < t; i++ {
		sum1 += arr[i][i]
		sum2 += arr[i][t-i-1]
	}
	return int32(math.Abs(float64(sum1) - float64(sum2)))
}
