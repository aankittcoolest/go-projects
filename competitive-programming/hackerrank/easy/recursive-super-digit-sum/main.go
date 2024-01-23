package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(superDigit("9875", 4))

}

func superDigit(n string, k int32) int32 {
	sum := digit(n)
	total := sum * int64(k)
	p := strconv.FormatInt(total, 10)
	for {
		res := digit(p)
		if int64(res/10) == 0 {
			return int32(res)
		}
		p = strconv.FormatInt(res, 10)
	}
}

func digit(p string) int64 {
	k := strings.Split(p, "")
	var sum int64
	for _, val := range k {
		item, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("something happened.")
		}
		sum += int64(item)
	}
	return sum
}
