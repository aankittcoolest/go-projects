package main

import "fmt"

func main() {
	//fmt.Println(solve(5, 2, 13))
	//fmt.Println(solve(6, 2, 13))
	fmt.Println(solve(38, 5, 81))

}

var arr = make([]int64, 1000001)

func solve(A int, B int, C int) int {
	D := factorial(int64(A), int64(C))
	E := factorial(int64(A-B), int64(C))
	F := factorial(int64(B), int64(C))
	fmt.Println(D % int64(C))
	fmt.Println(E % int64(C))
	fmt.Println(F % int64(C))
	fmt.Println(D % int64(C) / (E % int64(C) * F % int64(C)))
	return 1
	//return int((factorial(int64(A), int64(C))) / (factorial(int64(A-B), int64(C)) * factorial(int64(B), int64(C))))
}

func factorial(A int64, C int64) int64 {
	if A == 0 || A == 1 {
		return 1
	}
	return A * factorial(A-1, C)
	//return ((A % C) * (factorial(A-1, C) % C))
}
