package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
Add Two Numbers: https://leetcode.com/problems/add-two-numbers/
*/

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	number1 := getNumber(l1)
	fmt.Println(number1)

	number2 := getNumber(l2)
	fmt.Println(number2)

	sum := number1 + number2
	fmt.Println(sum)

	slice := strconv.Itoa(sum)
	fmt.Println(slice)
	sl := &ListNode{}
	for i := len(slice) - 1; i >= 0; i-- {
		fmt.Println(string(slice[i]))

		newNode := &ListNode{}
		newNode.Val, _ = strconv.Atoi(string(slice[i]))

		if i == len(slice)-1 {
			sl = newNode
		} else {
			curr := sl
			for curr.Next != nil {
				curr = curr.Next
			}
			curr.Next = newNode
		}
	}
	return sl
}

func getNumber(l *ListNode) int {
	var number int
	if l != nil {
		i := 1
		number = l.Val
		pow := 0
		for l.Next != nil {
			l = l.Next
			pow = int(math.Pow10(i))
			number = number + l.Val*pow
			i++
		}
	}
	return number
}

func main() {
	l1 := ListNode{Val: 7}
	l1.Next = &ListNode{Val: 1}

	l2 := ListNode{Val: 9}
	l2.Next = &ListNode{Val: 8}

	sl := addTwoNumbers(&l1, &l2)

	for sl != nil {
		fmt.Printf("digit: %d\n", sl.Val)
		sl = sl.Next
	}
}
