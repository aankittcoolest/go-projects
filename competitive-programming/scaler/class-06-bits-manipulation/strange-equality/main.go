package main

import (
	"fmt"
	"math"
	"strconv"
)

var debug bool = false

func main() {
	var answer int
	answer = solve(5)
	fmt.Println(answer)
}

func solve(A int) int {
	i := 0
	answer := 0
	for {
		if A>>i == 0 {
			break
		}
		bit := A >> i & 1
		if bit == 1 {
			answer = answer + int(math.Pow(float64(2), float64(i)))*0
		} else {
			answer = answer + int(math.Pow(float64(2), float64(i)))*1

		}
		i++
	}
	if debug {
		fmt.Println(answer)
		fmt.Println(strconv.FormatInt(int64(answer), 2))
	}

	return answer + (1 << i)

}
