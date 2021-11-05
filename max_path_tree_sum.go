package main

// https://leetcode.com/problems/binary-tree-maximum-path-sum/solution/
// O(h)
// I guess?
func maxPathSum(root *TreeNode) int {
	result := []int{root.Val}
	dfsMax(root, result)
	return result[0]
}

func dfsMax(root *TreeNode, result []int) int {
	if root == nil {
		return 0
	}
	leftMax := dfsMax(root.Left, result)
	rightMax := dfsMax(root.Right, result)
	leftMax = max(leftMax, 0)
	rightMax = max(rightMax, 0)

	// compute max path sum with split
	result[0] = max(result[0], root.Val+leftMax+rightMax)
	// without split
	return root.Val + max(leftMax, rightMax)
}
