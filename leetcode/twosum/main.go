package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
	fmt.Println(twoSum1(nums, 9))
	fmt.Println(twoSum2(nums, 9))

	nums = []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))
	fmt.Println(twoSum1(nums, 6))
	fmt.Println(twoSum2(nums, 6))

	nums = []int{3, 3}
	fmt.Println(twoSum(nums, 6))
	fmt.Println(twoSum1(nums, 6))
	fmt.Println(twoSum2(nums, 6))
}

// https://leetcode.com/problems/two-sum/solution/
// Brute Force appraoch with o(n*n)/o(1)
func twoSum(nums []int, target int) []int {
	for index, value := range nums {
		for after, val := range nums[index+1 : len(nums)] {
			if value+val == target {
				return []int{index, after + index + 1}
			}
		}
	}
	panic("no elements found")
}

// Two-pass Hash Table with o(n)/o(n)
func twoSum1(nums []int, target int) []int {
	indexValMap := map[int]int{}
	for index, value := range nums {
		indexValMap[value] = index
	}
	for index, value := range nums {
		found := indexValMap[target-value]
		if found > 0 && found != index {
			return []int{index, found}
		}
	}
	panic("no elements found")
}

// One-pass Hash Table with o(n)/o(n)
func twoSum2(nums []int, target int) []int {
	indexValMap := map[int]int{}
	for index, value := range nums {
		idx, found := indexValMap[target-value]
		if found {
			return []int{idx, index}
		}
		indexValMap[value] = index
	}
	panic("no elements found")
}
