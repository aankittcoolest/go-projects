package main

import (
	"fmt"
	"strings"
)

type stack []string

func (s stack) Push(a string) stack {
	return append(s, a)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	if l == 0 {
		return s, ""
	}

	return s[:l-1], s[l-1]
}

func main() {
	fmt.Println(isBalanced("[{()}]"))

}

func isBalanced(s string) string {
	a := strings.Split(s, "")

	var b stack

	for _, val := range a {
		if val == "{" || val == "[" || val == "(" {
			b = b.Push(val)
		} else {
			c := ""
			b, c = b.Pop()

			if (c == "(" && val == ")") || (c == "[" && val == "]") || (c == "{" && val == "}") {
				continue
			} else {
				return "NO"
			}
		}
	}

	if len(b) > 0 {
		return "NO"
	}

	return "YES"

}
