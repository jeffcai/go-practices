package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	rangeSumBST(testTree(), 1, 15)
}

func testTree() *TreeNode {
	tree := TreeNode{Val: 1, Left: nil, Right: nil}
	tree.Left = &TreeNode{Val: 5, Left: nil, Right: nil}
	tree.Right = &TreeNode{Val: 15, Left: nil, Right: nil}
	return &tree
}

/**
* Complexity
* time complexity: O(N)
* space complexity: O(1)
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	sum := 0
	fn := func(n *TreeNode) {
		if n != nil && n.Val >= L && n.Val <= R {
			sum += n.Val
		}
	}
	walk(root, fn)
	fmt.Println(sum)
	return sum
}

func walk(node *TreeNode, fn func(n *TreeNode)) {
	if node != nil {
		walk(node.Left, fn)
		fn(node)
		walk(node.Right, fn)
	}
}
