package main

import "fmt"

func main() {
	fmt.Println(pairs(2, []int32{1, 5, 3, 4, 2}))

}

func pairs(k int32, arr []int32) int32 {
	a := make(map[int32]int32)

	for _, val := range arr {
		a[val] = 1
	}

	var count int32
	for _, val := range arr {
		_, found := a[val+k]
		if found {
			count++
		}
	}

	return count

}
