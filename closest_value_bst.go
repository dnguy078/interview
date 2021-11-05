package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// Time: Avg O(log(n), worst case: O(n), tree that only has one branch
// Space: O(log(n))

func (tree *BST) FindClosestValue(target int) int {
	return tree.findClosest(target, tree.Value)
}

func (tree *BST) findClosest(target, closest int) int {
	if tree.Value == target {
		return target
	}
	if absolute(target, closest) > absolute(target, tree.Value) {
		closest = tree.Value
	}

	if target < tree.Value && tree.Left != nil {
		return tree.Left.findClosest(target, closest)
	} else if target > tree.Value && tree.Right != nil {
		return tree.Right.findClosest(target, closest)
	}
	return closest
}

func absolute(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
