package main

import (
	"fmt"
)

func main() {
	// solve(4, 5, []int{0, 1, 6, 11})
	// fmt.Println(solve(4, 5, []int{0, 1, 6, 11}))
	fmt.Println(solve(2, 100, []int{0, 1, 5}))
}

func enqueue(queue []int, element int) []int {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []int) (int, []int) {
	element := queue[0]
	if len(queue) == 1 {
		return element, []int{}
	}
	return element, queue[1:]
}

func solve(A int, B int, C []int) []int {
	queue := []int{}
	D := make([]int, 0)
	for i := 0; i < len(C); i++ {
		// fmt.Println(queue)
		for {
			if len(queue) > 0 {
				item := queue[0]
				if item <= C[i] {
					_, queue = dequeue(queue)
				} else {
					break
				}
			} else {
				break
			}
		}
		queue = enqueue(queue, B+C[i])
		if len(queue) > A {
			D = append(D, 0)
		} else {
			D = append(D, A-len(queue))
		}
	}
	return D
}
