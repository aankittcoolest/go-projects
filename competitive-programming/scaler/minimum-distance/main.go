package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(solve("x...o.x...o"))
	fmt.Println(solve("xo"))
	fmt.Println(solve("..x.o.."))
	fmt.Println(solve("xxx...xxx"))

}

func solve(A string) int {
	parts := strings.Split(A, "")
	current_count := 0
	min_count := math.MaxInt

	a := make([]string, 0)

	for i := 0; i < len(parts); i++ {
		if len(a) == 0 && (parts[i] == "x" || parts[i] == "o") {
			a = append(a, parts[i])
		} else if len(a) == 1 && a[0] == parts[i] {
			current_count = 0
		} else if len(a) == 1 && parts[i] == "." {
			current_count++
		} else if len(a) == 1 && parts[i] != a[0] && current_count > 0 {
			if min_count > current_count {
				min_count = current_count
			}
			a[0] = parts[i]
			current_count = 0
		}
	}
	if min_count == math.MaxInt {
		return -1
	}
	return min_count + 1
}
