package main

import "math"
import "fmt"

func main() {
	var answer int
	answer = solve([]int{1, 2, 3, 4, 5})
	fmt.Println(answer)

	answer = solve([]int{7, 8, 9, 10})
	fmt.Println(answer)

	answer = solve([]int{347148, 221001, 394957, 729925, 276769, 40726, 552988, 29952, 184491, 146773, 418965, 307, 219145, 183037, 178111, 81123, 109199, 683929, 422034, 346291, 11434, 7327, 340473, 316152, 364005, 359269, 170935, 105784, 224044, 22563, 48561, 165781, 9329, 357681, 169473, 175031, 605611, 374501, 6607, 329965, 76068, 836137, 103041, 486817, 195549, 107317, 34399, 56907, 37477, 189690, 36796, 376663, 39721, 177563, 174179, 183646, 217729, 408031, 429122, 631665, 282941, 526797, 262186, 306571, 63613, 57501, 70685, 226381, 1338, 9360, 130360, 20300, 400906, 87823, 180349, 108813, 18181, 119185, 1, 102611, 63591, 12889, 311185, 383896, 8701, 76077, 75481, 386017, 153553, 304913, 383455, 105948, 142885, 1, 12610, 137005, 119185, 16948, 66171, 123683})
	fmt.Println(answer)

}

func solve(A []int) int {
	total_count := (len(A) * (len(A) + 1)) / 2
	ans := 0
	for i := 0; i < 31; i++ {
		count := 0
		zeroes_count := 0
		for j := 0; j < len(A); j++ {
			if ((A[j] >> i) & 1) == 0 {
				zeroes_count++
			} else {
				if zeroes_count > 0 {
					count += (zeroes_count * (zeroes_count + 1)) / 2
					zeroes_count = 0
				}
			}
		}
		if zeroes_count > 0 {
			count += (zeroes_count * (zeroes_count + 1)) / 2
			zeroes_count = 0
		}

		ones_count := total_count - count
		ans = ans + ones_count*int(math.Pow(float64(2), float64(i)))
		ans = ans % 1000000007
	}
	return ans
}
