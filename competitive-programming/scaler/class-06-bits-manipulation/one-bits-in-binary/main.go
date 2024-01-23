package main

func main() {
	// numSetBits(11)
	numSetBits(6)

}

func numSetBits(A int) int {
	count := 0
	for {
		if A > 2 {
			if A%2 == 0 {
			} else {
				count++
			}
			A = A / 2
		} else if A == 2 || A == 1 {
			count++
			break
		} else {
			break
		}
	}
	return count
}
