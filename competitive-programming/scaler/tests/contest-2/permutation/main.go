package main

import "strings"
import "fmt"

func main() {
	fmt.Println(solve("0??"))
	fmt.Println(solve("?"))
	fmt.Println(solve("?1?"))

}

func solve(A string) int {
	B := strings.Split(A, "")
	ans := 1
	for i := 0; i < len(B); i++ {
		if i == 0 && B[i] == "0" {
			return 0
		}
		if i == 0 && B[i] == "?" {
			ans = ans * 9
		} else if B[i] == "?" {
			ans = ans * 10
		}

	}

	return ans
}
