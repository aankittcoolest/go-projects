package main

import (
	"fmt"
	"sort"
)

type item struct {
	value  int
	weight int
	profit float64
}

func main() {
	output := solve([]int{60, 100, 120}, []int{10, 20, 30}, 50)
	fmt.Println(output)
}

func solve(A []int, B []int, C int) int {
	items := make([]item, len(A))

	for i := 0; i < len(items); i++ {
		items[i] = item{
			A[i],
			B[i],
			float64(A[i] / B[i]),
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].profit > items[j].profit
	})
	total_value := 0
	for i := 0; i < len(items); i++ {
		if C >= items[i].weight {
			C = C - items[i].weight
			total_value += items[i].value
		} else {
			break
		}
	}
	return total_value
}
