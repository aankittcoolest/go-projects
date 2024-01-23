package main

import "fmt"

func main() {
	output := solve([]int{3, 34, 4, 12, 5, 2}, 9)
	fmt.Println(output)
	output = solve([]int{3, 34, 4, 12, 5, 2}, 30)
	fmt.Println(output)
	output = solve([]int {68, 35, 1, 70, 25, 79, 59, 63, 65, 6, 46}, 282)
	fmt.Println(output)
}


func solve(A []int, B int) bool {
	// create two-D array
	t := make([][]bool, len(A)+1)
	for i := 0; i < len(t); i++ {
		t[i] = make([]bool, B+1)
	}

	// fill in base condition
	for i:= 0; i<= len(A);i++ {
		for j:=0;j<= B;j++ {
			if(i== 0) {
				t[i][j] = false
			}
			if(j== 0) {
				t[i][j] = true
			}

		}
	}

	// approach problem by DP
	for i := 1; i <= len(A); i++ {
		for j := 1; j <= B; j++ {
			if A[i-1] <= j {
				t[i][j] = t[i-1][j-A[i-1]] || t[i-1][j]
			} else {
				t[i][j] = t[i-1][j]
			}
		}
	}
	return t[len(A)][B]

}
