package main

import (
	"sort"
)

type abc struct {
	item        string
	count       int
	moduleCount int
}

func main() {
	var a = []string{"Colorado", "Utah", "Wisconcin", "Oregon"}
	RemainderSorting(a)
}

func RemainderSorting(strArr []string) []string {
	var items []abc
	for _, val := range strArr {
		items = append(items, abc{val, len(val), len(val) % 3})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].moduleCount == items[j].moduleCount {
			a := 0
			b := 0
			for {
				if a < len(items[i].item) && b < len(items[j].item) {
					if items[i].item[a] < items[j].item[b] {
						return true
					} else if items[i].item[a] > items[j].item[b] {
						return false
					}
					a++
					b++
				} else {
					break
				}
			}
		}
		return items[i].moduleCount < items[j].moduleCount
	})
	var output []string
	for _, val := range items {
		output = append(output, val.item)
	}
	return output
}
