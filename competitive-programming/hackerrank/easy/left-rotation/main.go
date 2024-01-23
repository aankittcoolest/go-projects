package main

import "fmt"

func main() {
	fmt.Println(rotateLeft(2, []int32{1,2,3,4,5}))
}


func rotateLeft(d int32, arr []int32) []int32 {
	return append(arr[d:],arr[:d]...)
}
