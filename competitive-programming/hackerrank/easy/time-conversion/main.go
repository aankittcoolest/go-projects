package main

import (
	"fmt"
	"strconv"
)

func main() {
	//timeConversion("12:01:00AM")
	fmt.Println(timeConversion("12:00:00PM"))

}

func timeConversion(s string) string {
	hour := s[:2]
	minute := s[3:5]
	second := s[6:8]
	kind := s[8:10]

	h, err := strconv.Atoi(hour)
	if err != nil {
		fmt.Println("something happened.")
	}

	if kind == "AM" {
		if h >= 12 {
			h = 12 - h
		}

	} else if kind == "PM" {
		if h == 12 {
			h = 12
		} else {
			h = 12 + h
		}
	}
	output := fmt.Sprintf("%02d:%v:%v", h, minute, second)
	return output

}
