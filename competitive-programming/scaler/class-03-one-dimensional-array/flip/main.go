package main

import "fmt"
import "strings"
import "math"

func main() {
	var answer []int
	// answer = flip("010")
	// fmt.Println(answer)
	// answer = flip("111")
	// fmt.Println(answer)
	// answer = flip("1100")
	// fmt.Println(answer)
	// answer = flip("11001")
	// fmt.Println(answer)
	answer = flip("01010110")
	fmt.Println(answer)
}

func flip(A string) []int {
	res := strings.Split(A, "")
	count := 0
	maxCount := math.MinInt
	startIndex := -1
	endIndex := -1
	initalize := false
	for i := 0; i < len(res); i++ {
		if res[i] == "0" {
			initalize = true
			count++
		} else if res[i] == "1" && initalize {
			if maxCount < count {
				maxCount = count
				startIndex = i - count + 1
				endIndex = i
			}
			count = 0
		}
	}
	if count > 0 && maxCount < count {
		maxCount = count
		startIndex = len(res) - count + 1
		endIndex = len(res)
	}

	if startIndex == -1 {
		return []int{}
	}
	return []int{startIndex, endIndex}
}
