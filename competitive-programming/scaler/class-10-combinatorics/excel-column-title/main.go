package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(convertTitleToNumber("AB"))
	//fmt.Println(convertTitleToNumber("A"))
	//fmt.Println(convertTitleToNumber("BB"))
	//fmt.Println(convertTitleToNumber("HXR"))
	//fmt.Println(convertTitleToNumber("GJH"))
	//fmt.Println(convertToTitle(28))
	//fmt.Println(convertToTitle(1))
	//fmt.Println(convertToTitle(54))
	fmt.Println(convertToTitle(5000))

}

func convertTitleToNumber(A string) int {
	B := []rune(A)
	ans := 0
	j := 0
	for i := len(B) - 1; i >= 0; i-- {
		ans = ans + int(math.Pow(26, float64(j)))*(int(B[i])-64)
		j++
	}

	return ans
}

func convertToTitle(A int) string {
	ans := ""
	for {
		A = A - 1
		remainder := A % 26
		ans = string(65+remainder) + ans
		A = A / 26
		if A == 0 {
			break
		}
	}

	return ans
}
