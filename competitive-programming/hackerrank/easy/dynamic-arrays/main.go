package main

import "fmt"

func main() {
	fmt.Println(dynamicArray(2, [][]int32{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}))
}

func dynamicArray(n int32, queries [][]int32) []int32 {
	arr := make([][]int32, n)

	for i := 0; i < int(n); i++ {
		arr[i] = make([]int32, 0)
	}

	var lastAnswer int32
	lastAnswer = 0
	var lastAnswerArr []int32
	for i := range queries {
		ans := (queries[i][1] ^ lastAnswer) % n
		a := arr[ans]
		if queries[i][0] == 1 {
			a = append(a, queries[i][2])
			arr[ans] = a
		}
		if queries[i][0] == 2 {
			lastAnswer = a[queries[i][2]%int32(len(a))]
			lastAnswerArr = append(lastAnswerArr, lastAnswer)
		}
	}

	return lastAnswerArr
}
