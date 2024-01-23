// This solution is wrong and needs improvement

// credits to heap implementation
// https://www.sohamkamani.com/golang/heap/

package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

func main() {
	// ans := solve([][]int{
	// 	{1, 3},
	// 	{-2, 2},
	// }, 1)
	// fmt.Println(ans)

	// ans := solve([][]int{{40, 35}, {31, 40}, {20, 14}, {50, 45},
	// 	{48, 28}, {17, 44}, {19, 2}, {34, 41},
	// 	{44, 47}, {25, 14}, {44, 7}}, 11)
	// fmt.Println(ans)

	ans := solve([][]int{{48, 18}, {46, 34}, {47, 30}, {19, 9}, {1, 39},
		{95, 77}, {95, 106}, {78, 75},
		{92, 108}, {89, 83}, {80, 76}}, 5)
	fmt.Println(ans)
}

// we will impolement this solution via Heap
type Eucledian struct {
	distance int
	index    int
}

// for convenience let's implement string representation to print the Eucledian
func (e Eucledian) String() string {

	return "[" + strconv.Itoa(e.distance) + ", " + strconv.Itoa(e.index) + "]"
}

// creating new type to represent the list of Eucledian distances
type Eucledians []Eucledian

// implementing heap functions

func (e Eucledians) Less(d1, d2 int) bool {
	return d1 > d2
}

func (e Eucledians) Len() int {
	return len(e)
}

func (e Eucledians) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e *Eucledians) Push(x interface{}) {
	*e = append(*e, x.(Eucledian))
}

func (e *Eucledians) Pop() interface{} {
	old := *e
	n := len(old)
	x := old[n-1]
	*e = old[0 : n-1]
	return x
}

func getSquaredSum(x, y int) int {
	return x*x + y*y
}

func solve(A [][]int, B int) [][]int {
	eucledians := &Eucledians{}
	heap.Init(eucledians)
	for index, point := range A {
		heap.Push(eucledians, Eucledian{getSquaredSum(point[0], point[1]), index})
	}
	fmt.Println(eucledians)
	minimum_eucldedians := make([][]int, 0)
	for i := 0; i < B; i++ {
		minimum_eucldedian := eucledians.Pop()
		ans, ok := minimum_eucldedian.(Eucledian)
		if ok {
			minimum_eucldedians = append(minimum_eucldedians, A[ans.index])
		}

	}
	return minimum_eucldedians
}
