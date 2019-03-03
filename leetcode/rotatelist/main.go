package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	vals := []int{1, 2, 3}
	head := initList(vals)
	fmt.Println("head val: ", head.Val)

	newHead := rotateRight(head, 2)
	fmt.Println("new head val: ", newHead.Val, newHead.Next, newHead.Next.Next)
	len := lenOfListNodes(newHead)
	fmt.Println("length of new list: ", len)

	vals = []int{0, 1, 2}
	head = initList(vals)
	fmt.Println("head val: ", head.Val)

	newHead = rotateRight(head, 4)
	fmt.Println("new head val: ", newHead.Val, newHead.Next, newHead.Next.Next)
	len = lenOfListNodes(newHead)
	fmt.Println("length of new list: ", len)

	vals = []int{1}
	head = initList(vals)
	fmt.Println("head val: ", head.Val)

	newHead = rotateRight(head, 99)
	//	fmt.Println("new head val: ", newHead.Val, newHead.Next, newHead.Next.Next)
	len = lenOfListNodes(newHead)
	fmt.Println("length of new list: ", len)
}

func initList(vals []int) *ListNode {
	nodes := make([]ListNode, len(vals))
	for index := len(vals) - 1; index >= 0; index-- {
		nodes[index] = ListNode{Val: vals[index]}
		if index < len(vals)-1 {
			nodes[index].Next = &nodes[index+1]
		}
	}
	fmt.Println("lengh of nodes: ", len(nodes), len(vals))
	for _, node := range nodes {
		fmt.Println("node val: ", node.Val, node.Next)
	}
	return &nodes[0]
}

/*
 https://leetcode.com/problems/rotate-list/solution/
 time complexity:
 space complexity:
*/
func rotateRight(head *ListNode, k int) *ListNode {
	len := lenOfListNodes(head)

	if k == len || k == 0 || len == 0 || len == 1 || k%len == 0 {
		return head
	}

	if k > len {
		k = k % len
	}

	var newHead, newTail *ListNode
	node := head
	for index := 0; index < len; index++ {
		if index == len-k-1 {
			newTail = node
		}
		if index == len-k {
			newHead = node
		}
		if index == len-1 {
			node.Next = head
		}
		node = node.Next
	}

	// k <= len
	// k > len: k = k % len
	// len - k - 1 as tail, len - k as head => [len - k] ...  [len - 1] -> [0] ... [len - k - 1] -> nil
	newTail.Next = nil
	return newHead
}

func lenOfListNodes(head *ListNode) int {
	len := 0
	next := head
	for {
		if next != nil {
			next = next.Next
			len++
		} else {
			break
		}
	}
	fmt.Println("len of the list: ", len)
	return len
}
