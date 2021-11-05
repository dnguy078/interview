package main

// breadth first search visit itself then all its neighbors
// add itself to a list then its neighbors

func (n *Node) BreadthFirstSearch(array []string) []string {
	// Write your code here.
	queue := []*Node{n}
	var res []string
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		res = append(res, current.Name)
		queue = append(queue, current.Children...)
	}
	return res
}
