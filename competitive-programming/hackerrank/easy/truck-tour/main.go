package main

import "fmt"

// This is incomplete solution
// Test cases are still failing

func main() {
	fmt.Println(truckTour([][]int32{{1, 5}, {10, 3}, {12, 1}, {1, 1}, {3, 1}}))
}

func truckTour(petrolpumps [][]int32) int32 {
	var capacity int32 = 0
	var distance int32 = 0
	is_set := false

	var index int32 = -1
	for i, val := range petrolpumps {
		if distance+val[1] > capacity+val[0] {
			capacity = 0
			distance = 0
			is_set = false
		}
		if !is_set && val[0] >= val[1] {
			is_set = true
			index = int32(i)
			capacity += val[0]
			distance += val[1]
		}
	}
	return index
}
