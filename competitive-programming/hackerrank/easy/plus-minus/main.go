package main

import "fmt"

func main() {
	plusMinus([]int32{1, 1, 0, -1, -1})

}

func plusMinus(arr []int32) {
	total := float64(len(arr))
	var plusCount, minusCount, zeroCount int32 = 0, 0, 0
	for _, val := range arr {
		if val > 0 {
			plusCount++
			continue
		}
		if val < 0 {
			minusCount++
			continue
		} else {
			zeroCount++
		}
	}
	fmt.Printf("%.6f\n", float64(plusCount)/total)
	fmt.Printf("%.6f\n", float64(minusCount)/total)
	fmt.Printf("%.6f\n", float64(zeroCount)/total)
}
