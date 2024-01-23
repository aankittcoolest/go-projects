package main

import (
	"fmt"
	"sort"
)

func main() {
	ans := gridChallenge([]string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"})
	fmt.Println(ans)
	//fmt.Println([]string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"})
}

func gridChallenge(grid []string) string {
	var r [][]rune
	for _, val := range grid {
		a := []rune(val)
		sort.Slice(a, func(i, j int) bool {
			return a[i] < a[j]
		})
		r = append(r, a)
	}

	for i := 0; i < len(r)-1; i++ {
		for k := 0; k < len(r[i]); k++ {
			if r[i+1][k] < r[i][k] {
				return "NO"
			}
		}
	}
	return "YES"

	// Write your code here

}
