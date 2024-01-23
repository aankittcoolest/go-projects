package main

import "fmt"

func main() {
	cpFact(30, 12)
	cpFact(5, 10)

}

func cpFact(A int, B int) int {
	max := -1
	for i := 1; i*i < A; i++ {
		if i == A/i {
			if gcd(i, B) == 1 {
				if i > max {
					max = i
				}
			}
		}
		if i*(A/i) == A {
			if gcd(i, B) == 1 {
				if i > max {
					max = i
				}
			}
			if gcd(A/i, B) == 1 {
				if A/i > max {
					max = A / i
				}
			}
		}
	}
	fmt.Println(max)
	return max
}

func gcd(A int, B int) int {
	if B == 0 {
		return A
	}
	C := A % B
	return gcd(B, C)
}
