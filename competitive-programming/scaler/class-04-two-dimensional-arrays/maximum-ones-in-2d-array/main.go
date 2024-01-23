package main

import (
	"fmt"
	"math"
)

func main() {
	var answer int
	answer = solve([][]int{
		{0, 1, 1},
		{0, 0, 1},
		{0, 1, 1},
	})
	fmt.Println(answer)

	answer = solve([][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 1, 1},
		{0, 1, 1, 1},
	})
	fmt.Println(answer)

}

func solve(A [][]int) int {

	i := 0
	j := len(A[i]) - 1

	max_count := math.MinInt
	count := 0
	index := 0
	for {
		if i == len(A) || j == -1 {
			if max_count < count {
				index = i
				max_count = count
			}
			break
		}

		if A[i][j] == 0 {
			if max_count < count {
				index = i
				max_count = count
			}
			i++
			continue
		}
		if A[i][j] == 1 {
			count++
			j--
			continue
		}

	}

	return index

}
