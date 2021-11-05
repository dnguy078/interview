package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	list := make([]int, 0)
	preorder(root, list)
	return list
}

func preorder(node *TreeNode, input []int) {
	for node != nil {
		input = append(input, node.Val)
		preorder(node.Left, input)
		preorder(node.Right, input)
	}
}
