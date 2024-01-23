package main

import "fmt"

func main() {
	// answer := solve([]int{1, 2, 3, 1, 2, 4})
	// fmt.Println(answer)

	answer := solve([]int{684, 96, 1126, 1288, 330, 1360, 940, 330, 1288, 580, 476, 812, 96, 264, 1360, 684, 476, 1126, 580, 812})
	fmt.Println(answer)

}

func solve(A []int) []int {
	B := make([]int, 0)
	C := make([]int, 0)
	for i := 0; i < 31; i++ {
		set_count := 0
		unset_count := 0
		for j := 0; j < len(A); j++ {
			if A[j]>>i&1 == 1 {
				set_count++
			} else {
				unset_count++
			}
		}
		if set_count%2 != 0 {
			for j := 0; j < len(A); j++ {
				if ((A[j] >> i) & 1) == 1 {
					B = append(B, A[j])
					set_count++
				} else {
					C = append(C, A[j])
					unset_count++
				}
			}
			break
		}

	}
	b := B[0]
	c := C[0]

	for i := 1; i < len(B); i++ {
		b = b ^ B[i]
	}
	for i := 1; i < len(C); i++ {
		c = c ^ C[i]
	}

	if b > c {
		return []int{c, b}
	}
	return []int{b, c}

}
