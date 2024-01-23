package main

import "fmt"

var count = 0

func main() {
	//expected answer 4
	abc(0, 0, 1, 3)
	fmt.Println(count)
}

func abc(sr int, sc int, dr int, dc int) {
	if sr > dr || sc > dc {
		return
	}

	if sr == dr && sc == dc {
		count++
		return
	}

	abc(sr, sc+1, dr, dc)
	abc(sr+1, sc, dr, dc)
}
