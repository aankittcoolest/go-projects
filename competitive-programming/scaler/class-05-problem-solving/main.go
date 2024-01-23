package main

import (
	"fmt"
	"math"
)

type Interval struct {
	start, end int
}

func main() {
	test_case_3()
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	answer := make([]Interval, 0)
	for i := 0; i < len(intervals); i++ {
		if intervals[i].end < newInterval.start {
			answer = append(answer, intervals[i])
		} else if intervals[i].start > newInterval.end {
			answer = append(answer, newInterval)
			for j := i; j < len(intervals); j++ {
				answer = append(answer, intervals[j])
			}
			return answer
		} else {
			newInterval.start = int(math.Min(float64(intervals[i].start), float64(newInterval.start)))
			newInterval.end = int(math.Max(float64(intervals[i].end), float64(newInterval.end)))
		}
	}
	answer = append(answer, newInterval)
	return answer
}

func test_case_3() {
	intervals := make([]Interval, 0)
	intervals = append(intervals, Interval{1, 2})
	intervals = append(intervals, Interval{3, 4})
	intervals = append(intervals, Interval{5, 6})

	newInterval := Interval{4, 5}
	answer := insert(intervals, newInterval)
	fmt.Println(answer)
}

func test_case_2() {
	intervals := make([]Interval, 0)
	intervals = append(intervals, Interval{1, 3})
	intervals = append(intervals, Interval{4, 7})
	intervals = append(intervals, Interval{10, 14})
	intervals = append(intervals, Interval{16, 19})
	intervals = append(intervals, Interval{21, 24})
	intervals = append(intervals, Interval{27, 30})
	intervals = append(intervals, Interval{32, 35})

	newInterval := Interval{10, 22}
	answer := insert(intervals, newInterval)
	fmt.Println(answer)
}

func test_case_1() {
	intervals := make([]Interval, 0)
	intervals = append(intervals, Interval{6037774, 6198243})
	intervals = append(intervals, Interval{6726550, 7004541})
	intervals = append(intervals, Interval{8842554, 10866536})
	intervals = append(intervals, Interval{11027721, 11341296})
	intervals = append(intervals, Interval{11972532, 14746848})
	intervals = append(intervals, Interval{16374805, 16706396})
	intervals = append(intervals, Interval{17557262, 20518214})
	intervals = append(intervals, Interval{22139780, 22379559})
	intervals = append(intervals, Interval{27212352, 28404611})
	intervals = append(intervals, Interval{28921768, 29621583})
	intervals = append(intervals, Interval{29823256, 32060921})
	intervals = append(intervals, Interval{33950165, 36418956})
	intervals = append(intervals, Interval{37225039, 37785557})
	intervals = append(intervals, Interval{40087908, 41184444})
	intervals = append(intervals, Interval{41922814, 45297414})
	intervals = append(intervals, Interval{48142402, 48244133})
	intervals = append(intervals, Interval{48622983, 50443163})
	intervals = append(intervals, Interval{50898369, 55612831})
	intervals = append(intervals, Interval{57030757, 58120901})
	intervals = append(intervals, Interval{59772759, 59943999})
	intervals = append(intervals, Interval{61141939, 64859907})
	intervals = append(intervals, Interval{65277782, 65296274})
	intervals = append(intervals, Interval{67497842, 68386607})
	intervals = append(intervals, Interval{70414085, 73339545})
	intervals = append(intervals, Interval{73896106, 75605861})
	intervals = append(intervals, Interval{79672668, 84539434})
	intervals = append(intervals, Interval{84821550, 86558001})
	intervals = append(intervals, Interval{91116470, 92198054})
	intervals = append(intervals, Interval{96147808, 98979097})

	newInterval := Interval{36210193, 61984219}
	answer := insert(intervals, newInterval)
	fmt.Println(answer)

}

// (6037774, 6198243) (6726550, 7004541) (8842554, 10866536) (11027721, 11341296) (11972532, 14746848) (16374805, 16706396) (17557262, 20518214) (22139780, 22379559) (27212352, 28404611) (28921768, 29621583) (29823256, 32060921) (33950165, 64859907) (65277782, 65296274) (67497842, 68386607) (70414085, 73339545) (73896106, 75605861) (79672668, 84539434) (84821550, 86558001) (91116470, 92198054) (96147808, 98979097)
