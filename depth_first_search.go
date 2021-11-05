package main

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	for _, c := range n.Children {
		array = c.DepthFirstSearch(array)
	}
	return array
}
