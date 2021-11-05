package main

// type BinaryTree struct {
// 	Value       int
// 	Left, Right *BinaryTree
// }

func NodeDepths(root *BinaryTree) int {
	return sumNodeDepths(root, 0)
}

func sumNodeDepths(node *BinaryTree, depth int) int {
	if node == nil {
		return 0
	}
	return depth + sumNodeDepths(node.Left, depth+1) + sumNodeDepths(node.Right, depth+1)
}
