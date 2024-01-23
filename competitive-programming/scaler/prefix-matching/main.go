package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(solve([]string{"a", "aa", "aab", "b", "bb", "bba", "bbb"}, "bb"))
	fmt.Println(solve([]string{"a", "b"}, "c"))
}

func solve(A []string, B string) []int {
	r, _ := regexp.Compile("^" + B)
	start_index := -1
	end_index := -1
	for i := 0; i < len(A); i++ {
		result := r.MatchString(A[i])
		if result {
			if start_index == -1 {
				start_index = i
			} else {
				end_index = i
			}
		} else if start_index != -1 && !result {
			break
		}
	}
	return []int{start_index, end_index}
}
