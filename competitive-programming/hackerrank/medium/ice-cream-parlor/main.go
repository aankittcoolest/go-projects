package main

import "fmt"

func main() {
	var a = []int32{2, 1, 3, 5, 6}
	whatFlavors(a, 5)

}

func whatFlavors(cost []int32, money int32) {
	var b = make(map[int32]int32)
	for i, val := range cost {
		remaining := money - val
		item, found := b[remaining]
		if found {
			fmt.Println(item+1, i+1)
			break
		}
		b[val] = int32(i)
	}
}
