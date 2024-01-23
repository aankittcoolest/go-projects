package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	start, end int
}

func main() {
	test_case_1()
	test_case_2()
	test_case_3()
}

func merge(A []Interval) []Interval {
	sort.Slice(A, func(i, j int) bool {
		return A[i].start < A[j].start
	})
	i := 0
	B := make([]Interval, 0)
	min := A[0].start
	max := A[0].end
	for {
		if i+1 == len(A) {
			break
		} else if max >= A[i+1].start {
			min = int(math.Min(float64(min), float64(A[i+1].start)))
			max = int(math.Max(float64(max), float64(A[i+1].end)))
		} else {
			B = append(B, Interval{min, max})
			min = A[i+1].start
			max = A[i+1].end
		}
		i++
	}
	B = append(B, Interval{min, max})

	return B
}

func test_case_1() {
	A := make([]Interval, 0)
	A = append(A, Interval{2, 6})
	A = append(A, Interval{1, 3})
	A = append(A, Interval{8, 10})
	A = append(A, Interval{15, 18})
	answer := merge(A)
	fmt.Println(answer)
}

func test_case_2() {
	// 6 1 10 2 9 3 8 4 7 5 6 6 6
	A := make([]Interval, 0)
	A = append(A, Interval{1, 10})
	A = append(A, Interval{2, 9})
	A = append(A, Interval{3, 8})
	A = append(A, Interval{4, 7})
	A = append(A, Interval{5, 6})
	A = append(A, Interval{6, 6})
	answer := merge(A)
	fmt.Println(answer)

}

func test_case_3() {
	abc := "89 47 76 51 99 28 78 54 94 12 72 31 72 12 55 24 40 59 79 41 100 46 99 5 27 13 23 9 69 39 75 51 53 81 98 14 14 27 89 73 78 28 35 19 30 39 87 60 94 71 90 9 44 56 79 58 70 25 76 18 46 14 96 43 95 70 77 13 59 52 91 47 56 63 67 28 39 51 92 30 66 4 4 29 92 58 90 6 20 31 93 52 75 41 41 64 89 64 66 24 90 25 46 39 49 15 99 50 99 9 34 58 96 85 86 13 68 45 57 55 56 60 74 89 98 23 79 16 59 56 57 54 85 16 29 72 86 10 45 6 25 19 55 21 21 17 83 49 86 67 84 8 48 63 85 5 31 43 48 57 62 22 68 48 92 36 77 27 63 39 83 38 54 31 69 36 65 52 68"
	B := strings.Split(abc, " ")

	A := make([]Interval, 0)
	for i := 1; i < len(B); i = i + 2 {
		a, _ := strconv.Atoi(B[i])
		b, _ := strconv.Atoi(B[i+1])
		A = append(A, Interval{a, b})
	}
	answer := merge(A)
	fmt.Println(answer)

}
