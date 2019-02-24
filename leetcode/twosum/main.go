package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))

	nums = []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))

}

//https://leetcode.com/problems/two-sum/solution/
// Brute Force appraoch with o(n*n)/o(1)
func twoSum(nums []int, target int) []int {
	for index, value := range nums {
		for after, val := range nums[index+1 : len(nums)] {
			fmt.Println(strconv.Itoa(index) + " " + strconv.Itoa(after))
			if value+val == target {
				return []int{index, after + index + 1}
			}
		}
	}
	return []int{0, 0}
}
