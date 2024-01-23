package main

import "fmt"

func main() {
	output := solve([]int {1,2,2,1}, []int {2,3,1,2})
	fmt.Println(output)
	output = solve([]int {2,1,4,10}, []int {3,6,2,10,10})
	fmt.Println(output)
}

func solve(A []int , B []int )  ([]int) {
	 C:= make([]int, 0 )
	 for i:=0; i<len(A);i++ {
		for j:=0; j<len(B);j++ {
			if A[i] == B[j] {
				C = append(C, A[i])
				B = append(B[:j],B[j+1:]...)
				break
			}
		}
	 }
	 return C
}

