package main

import (
	"fmt"
	"math"
	"sort"
	//	"strconv"
)

func main() {
	var answer int
	answer = findMinXor([]int{0, 2, 5, 7})
	fmt.Println(answer)
	answer = findMinXor([]int{10, 23, 16, 19, 14})
	fmt.Println(answer)
	answer = findMinXor([]int{4, 13, 6, 7, 10})
	fmt.Println(answer)
	answer = findMinXor([]int{1164510, 462942, 361856, 1278046, 1128414, 789726, 35328, 487424, 983390, 1043550, 634624, 109440, 1247326, 944384, 337758, 161920, 908032, 686174, 1178334, 365022, 1275392, 449246, 448384, 700510, 464640, 708864, 716510, 857054, 226816, 946910, 332416, 1178624, 1161088, 1172608, 896862, 374016, 256478, 756608, 503646, 452864, 651008, 925534, 951040, 1194974, 663040, 695168, 91742, 402654, 1106432, 471646, 1225950, 136542, 951646, 653918, 1062622, 1246720, 1210496, 871168, 171614, 155776, 319070, 848862, 614784, 156254, 603102, 218846, 257152, 182016, 849536, 964224, 625374, 805376, 1001344, 650334, 1055488, 508928})
	fmt.Println(answer)
}

func findMinXor(A []int) int {
	min_output := math.MaxInt
	sort.Ints(A)
	for j := 0; j < len(A)-1; j++ {
		output := A[j] ^ A[j+1]
		if output < min_output {
			min_output = output
		}
	}
	return min_output
}
