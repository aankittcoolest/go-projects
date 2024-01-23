package main

import "fmt"

func main() {
	var answer int
	answer = solve([]int{1, 2, 3, 4, 5}, 2)
	fmt.Println(answer)
	answer = solve([]int{5, 17, 100, 11}, 28)
	fmt.Println(answer)
}

func solve(A []int, B int) int {
	C := make([]int, B)
	for i := 0; i < len(A); i++ {
		remainder := A[i] % B
		C[remainder]++
	}
	answer := 0
	answer = (answer + C[0]*(C[0]-1)/2) % 1000000007
	if B%2 == 0 {
		answer = (answer + C[B/2]*(C[B/2]-1)/2) % 1000000007
	}

	i := 1
	j := len(C) - 1
	for {
		if i < j {
			answer = (answer + C[i]*C[j]) % 1000000007
			i++
			j--
		} else {
			break
		}
	}
	return answer
}
