package main

import "fmt"

func main() {
	var answer []int
	answer = solve(5, [][]int{
		{1, 2, 10},
		{2, 3, 20},
		{2, 5, 25},
	})
	fmt.Println(answer)

}

func solve(A int, B [][]int) []int {
	C := make([]int, A)

	for i := 0; i < len(B); i++ {
		C[B[i][0]-1] = C[B[i][0]-1] + B[i][2]
		if B[i][1]+1 < len(C)+1 {
			C[B[i][1]] = C[B[i][1]] - B[i][2]
		}
	}

	//create psum array
	for i := 1; i < len(C); i++ {
		C[i] = C[i] + C[i-1]
	}
	return C
}
