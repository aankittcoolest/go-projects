package main

import "fmt"

func main() {
	fmt.Println(lonelyinteger([]int32{1, 2, 3, 4, 3, 2, 1}))

}

func lonelyinteger(a []int32) int32 {
	// Write your code here
	m := make(map[int32]int32)
	for _, val := range a {
		item, found := m[val]
		if found {
			m[val] = item + 1
		} else {
			m[val] = 1
		}
	}

	for key, val := range m {
		if val%2 != 0 {
			return key
		}
	}
	return 1
}
