package main

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	var sums []int
	branchSum(root, 0, &sums)
	return sums
}

func branchSum(node *BinaryTree, sum int, sums *[]int) {
	if node == nil {
		return
	}
	runningSum := sum + node.Value
	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, runningSum)
		return
	}
	branchSum(node.Left, runningSum, sums)
	branchSum(node.Right, runningSum, sums)
}
