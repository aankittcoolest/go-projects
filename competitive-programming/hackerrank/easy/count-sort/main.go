package main

import "fmt"

func main() {
	fmt.Println(countingSort([]int32{1,1,3,2,1}))

}

func countingSort(arr []int32) []int32 {
	var result []int32

	for i:=0;i<100;i++ {
		result = append(result,int32(0))
	}

	for _,val := range arr {
		result[val] = result[val] + 1
	}
	return result
}