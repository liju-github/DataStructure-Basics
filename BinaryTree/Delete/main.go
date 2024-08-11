package main

type Node struct{
	Value int
	Left *Node
	Right *Node
}

func (n *Node) Delete(value int) *Node {
	if n == nil {
		return nil
	}

	// Traverse to find the node to delete
	if value < n.Value {
		n.Left = n.Left.Delete(value)
	} else if value > n.Value {
		n.Right = n.Right.Delete(value)
	} else {
		// Node found, handle the three cases

		// Case 1: No children (leaf node)
		if n.Left == nil && n.Right == nil {
			return nil
		}

		// Case 2: One child
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		// Case 3: Two children
		// Find the in-order successor (smallest in the right subtree)
		smallestRight := n.Right.findMin()
		n.Value = smallestRight.Value
		n.Right = n.Right.Delete(smallestRight.Value)
	}

	return n
}

// Helper function to find the minimum value node in a subtree
func (n *Node) findMin() *Node {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}
