package main

import "fmt"

func main() {
	fmt.Println("good")
	fmt.Println(utopianTree(0))
}

func utopianTree(n int32) int32 {
	// Write your code here
	var height int32
	for i := 0; i <= int(n); i++ {
		if i%2 == 0 {
			height = height + 1
		} else {
			height = height * 2
		}

	}
	return height
}
