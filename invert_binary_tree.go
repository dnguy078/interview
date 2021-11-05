package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left, right := invertTree(root.Left), invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
