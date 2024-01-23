package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMedian([]int32{5, 3, 1, 2, 4}))

}

func findMedian(arr []int32) int32 {
	var newArr []int
	for _, val := range arr {
		newArr = append(newArr, int(val))
	}
	sort.Ints(newArr)
	mid := len(newArr) / 2
	return int32(newArr[mid])

}
