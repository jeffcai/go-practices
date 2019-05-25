package main

import (
	"fmt"
)

func main() {

	fmt.Println(numJewelsInStones("az", "azBBA"))
	fmt.Println(numJewelsInStones("CC", "azBBA"))
	fmt.Println(numJewelsInStones("aszB", "azBBA"))
}

/*
 https://leetcode.com/problems/jewels-and-stones
 time complexity: O(N)
 space complexity: O(1)
*/
func numJewelsInStones(J string, S string) int {
	jewels := []byte(J)
	jewelsMap := toMap(jewels)

	stones := []byte(S)

	count := 0
	for _, stone := range stones {
		if val := jewelsMap[stone]; val > 0 {
			count++
		}
	}
	return count
}

func toMap(items []byte) map[byte]byte {
	set := make(map[byte]byte, len(items))
	for _, item := range items {
		set[item] = item
	}
	return set
}
