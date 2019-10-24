package main

import (
	"fmt"
)

// The problem: https://leetcode.com/problems/smallest-integer-divisible-by-k/submissions/
// Refer to this solution: https://leetcode.com/problems/smallest-integer-divisible-by-k/discuss/309846/Go-solution
func smallestRepunitDivByK(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}

	r := 0
	// why this loop works?
	for N := 1; N <= K; N++ {
		r = (r*10 + 1) % K
		if r == 0 {
			return N
		}
	}

	return -1
}

func main() {
	n := smallestRepunitDivByK(1)
	fmt.Printf("smallest %d \n", n)

	n = smallestRepunitDivByK(2)
	fmt.Printf("smallest %d \n", n)

	n = smallestRepunitDivByK(23)
	fmt.Printf("smallest %d \n", n)
}
