package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	lc := lowestCommonAncestor(root.Left, p, q)
	rc := lowestCommonAncestor(root.Right, p, q)
	if lc == nil {
		return rc
	}
	if rc == nil {
		return lc
	}
	return root
}

func lowestCommonAncestorInBST(root, p, q *TreeNode) *TreeNode {
	cur := root

	for cur != nil {
		if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else {
			return cur
		}
	}
	return cur

}
