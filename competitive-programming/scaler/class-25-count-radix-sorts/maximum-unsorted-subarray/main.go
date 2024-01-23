package main

import "fmt"

func main() {
	// subUnsort([]int{1, 2, 5, 3, 4, 7, 6, 8})
	fmt.Println(subUnsort([]int{1, 1}))
	fmt.Println(subUnsort([]int{2, 1, 1}))
	// fmt.Println(subUnsort([]int{1,1,10,10,15,10,15,10,10,15,10,15}))
	// fmt.Println(subUnsort([]int{4, 15, 4, 4, 15, 18, 20}))
	// fmt.Println(subUnsort([]int{3, 3, 4, 5, 5, 9, 11, 13, 14, 15, 15, 16, 15, 20, 16}))

}

func subUnsort(A []int) []int {
	if len(A) <= 1 {
		return []int{-1}
	}
	min_index := -1
	max_index := -1

	final_min_index := -1
	final_max_index := -1

	for i := 1; i < len(A); i++ {
		if A[i] < A[i-1] {
			min_index = i - 1
			final_min_index = min_index
			val := A[min_index]

			for j := min_index - 1; j >= 0; j-- {
				if A[j] == val {
					final_min_index = j
					continue
				} else {
					break
				}
			}
			break
		}
	}
	if final_min_index == -1 {
		return []int{-1}
	}
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] > A[i+1] {
			max_index = i + 1
			final_max_index = max_index
			val := A[max_index]

			for j := max_index + 1; j <= len(A)-1; j++ {
				if A[j] == val {
					final_max_index = j
					continue
				} else {
					break
				}
			}
			break
		}
	}

	return []int{final_min_index, final_max_index}
	// for i := final_min_index; i < final_max_index; i++ {
	// 	if A[i] == A[i+1] {
	// 		continue
	// 	} else {
	// 		return []int{final_min_index, final_max_index}

	// 	}
	// }

	// // handle edge case {1,1}
	// return []int{-1}
}

// [4,15,4,4,15,18,20]
