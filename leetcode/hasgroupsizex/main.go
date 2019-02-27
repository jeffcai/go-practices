package main

import (
	"fmt"
)

func main() {
	has := hasGroupsSizeX([]int{1, 1})
	fmt.Println("hax groups size x: t% ", has)

	has = hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2})
	fmt.Println("hax groups size x: t% ", has)

	has = hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2})
	fmt.Println("hax groups size x: t% ", has)

	has = hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3})
	fmt.Println("hax groups size x: t% ", has)

	has = hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1})
	fmt.Println("hax groups size x: t% ", has)

	has = hasGroupsSizeXWithGCD([]int{1, 1})
	fmt.Println("hax groups size x: t% ", has)

	has = hasGroupsSizeXWithGCD([]int{1, 1, 2, 2, 2, 2})
	fmt.Println("hax groups size x: t% ", has)

	has = hasGroupsSizeXWithGCD([]int{1, 1, 1, 2, 2, 2})
	fmt.Println("hax groups size x: t% ", has)

	has = hasGroupsSizeXWithGCD([]int{1, 1, 1, 2, 2, 2, 3, 3})
	fmt.Println("hax groups size x: t% ", has)

	has = hasGroupsSizeXWithGCD([]int{1, 2, 3, 4, 4, 3, 2, 1})
	fmt.Println("hax groups size x: t% ", has)
}

// https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/solution/
// Brute Force
// Time Complexity:
// Space Complexity: O(N)
func hasGroupsSizeX(deck []int) bool {
	num := len(deck)
	if num >= 2 {
		//		fmt.Println("num is " + strconv.Itoa(num))
		for factor := 2; factor <= num; factor++ {
			if num%factor == 0 {
				usedMap := map[int]bool{}
				for i, value := range deck {
					_, used := usedMap[i]
					if !used {
						count := 1
						for j := i + 1; j < num; j++ {
							if deck[j] == value {
								usedMap[j] = true
								count++
							}
							if count == factor {
								usedMap[i] = true
								//								fmt.Println("meet current factor: %d", factor)
								break
							}
						}
					}
					//					fmt.Println("current used map len: %d, %d", i, len(usedMap))
					if i == num-1 && len(usedMap) == num {
						//						fmt.Println("factor is " + strconv.Itoa(factor))
						return true
					}
				}
			}
		}
	}
	return false
}

// Greatest Common Divisor
// Time Complexity:
// Space Complexity: O(N)
func hasGroupsSizeXWithGCD(deck []int) bool {
	counts := [10000]int{}
	for _, value := range deck {
		counts[value]++
	}
	g := -1
	for _, count := range counts {
		if g == -1 {
			g = count
		} else {
			g = GCDRemainderRecursive(g, count)

		}
	}
	return g >= 2
}

// GCDEuclidean calculates GCD by Euclidian algorithm.
func GCDEuclidean(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

// GCDRemainderRecursive calculates GCD recursively using remainder.
func GCDRemainderRecursive(a, b int) int {
	if b == 0 {
		return a
	}
	return GCDRemainderRecursive(b, a%b)
}

// GCDRemainder calculates GCD iteratively using remainder.
func GCDRemainder(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
