package main

import (
	"fmt"
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
	head := &ListNode{}
	p, q, curr := l1, l2, head
	var carry int
	for p != nil || q != nil {
		var x, y int
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		sum := x + y + carry
		fmt.Println(sum)
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return head.Next
}

//func getNumber(l *ListNode) int {
//	var number int
//	if l != nil {
//		i := 1
//		number = l.Val
//		pow := 0
//		for l.Next != nil {
//			l = l.Next
//			pow = int(math.Pow10(i))
//			number = number + l.Val*pow
//			i++
//		}
//	}
//	return number
//}

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
