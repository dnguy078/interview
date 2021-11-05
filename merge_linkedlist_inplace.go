package main

// This is an input struct. Do not edit.
type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MergeLinkedLists(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	p1 := headOne
	p2 := headTwo
	var prev *LinkedList
	for p1 != nil && p2 != nil {
		if p1.Value < p2.Value {
			prev = p1
			p1 = p1.Next
		} else {
			if prev != nil {
				prev.Next = p2
			}
			prev = p2
			p2 = prev.Next
			prev.Next = p1
		}
	}
	if p1 == nil {
		prev.Next = p2
	}

	if headOne.Value < headTwo.Value {
		return headOne
	}
	return headTwo
}
