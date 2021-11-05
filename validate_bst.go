package main

import (
	"math"
)

// version from leetcode
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	if node.Left != nil && !validate(node.Left, min, node.Val) {
		return false
	}
	if node.Right != nil && !validate(node.Right, node.Val, max) {
		return false
	}
	return true
}

/// version two from algo expert
// type BST struct {
// 	Value int

// 	Left  *BST
// 	Right *BST
// }

func (tree *BST) ValidateBst() bool {
	return tree.validate(-100000, 1000000)
}

func (tree *BST) validate(min, max int) bool {
	if tree.Value < min || tree.Value >= max {
		return false
	}
	if tree.Left != nil && !tree.Left.validate(min, tree.Value) {
		return false
	}
	if tree.Right != nil && !tree.Right.validate(tree.Value, max) {
		return false
	}
	return true
}
