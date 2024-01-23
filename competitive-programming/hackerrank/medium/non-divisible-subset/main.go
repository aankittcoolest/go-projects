package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("good")
	nonDivisibleSubset(7, []int32{278, 576, 496, 727, 410, 124, 338, 149, 209, 702, 282, 718, 771, 575, 436})
	nonDivisibleSubset(3, []int32{1, 7, 2, 4})
}

func nonDivisibleSubset(k int32, s []int32) int32 {
	m := make(map[int32]int32)
	for i := 0; i < len(s); i++ {
		m[s[i]%k]++
	}

	var ans int32
	if m[0] > 0 {
		ans++
	}

	for i := 1; i < int(k); i++ {
		if m[int32(i)]%k == 0 {
			continue
		}

		if i+i == int(k) {
			ans++
		} else {
			ans += int32(math.Max(float64(m[int32(i)]), float64(m[k-int32(i)])))
			m[int32(i)] = 0
			m[k-int32(i)] = 0
		}

	}
	return ans
}
