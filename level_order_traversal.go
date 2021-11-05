package main

// https://leetcode.com/problems/binary-tree-level-order-traversal/submissions/
func levelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	queue = append(queue, root)
	var results [][]int

	for len(queue) != 0 {
		qLen := len(queue)
		var level []int
		for i := 0; i < qLen; i++ {
			n := queue[0]
			queue = queue[1:]
			if n != nil {
				level = append(level, n.Val)
				queue = append(queue, n.Left)
				queue = append(queue, n.Right)
			}
		}
		if len(level) != 0 {
			results = append(results, level)
		}
	}
	return results
}
