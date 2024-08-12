package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Delete(value int) {
	if n == nil {
		return
	}
	queue := []*Node{n}
	var targetNode, deepestNode *Node
	var parentOfDeepest *Node

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Value == value {
			targetNode = current
		}
		if current.Left != nil {
			parentOfDeepest = current
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			parentOfDeepest = current
			queue = append(queue, current.Right)
		}
		deepestNode = current
	}

	if targetNode != nil {
		targetNode.Value = deepestNode.Value
		if parentOfDeepest.Right == deepestNode {
			parentOfDeepest.Right = nil
		} else {
			parentOfDeepest.Left = nil
		}
	}
}

func (n *Node) Insert(value int) {

	newnode := &Node{Value: value}
	if n == nil {
		n = newnode
	}

	queue := []*Node{n}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Left == nil {
			current.Left = newnode
			break
		} else {
			queue = append(queue, current.Left)
		}

		if current.Right == nil {
			current.Right = newnode
			break
		} else {
			queue = append(queue, current.Right)
		}
	}
}

func (n *Node) Inorder() {
	if n == nil {
		return
	}

	n.Left.Inorder()
	fmt.Print("->", n.Value)
	n.Right.Inorder()
}

func main() {
	n := &Node{Value: 100}

	n.Insert(10)
	n.Insert(30)
	n.Insert(22)
	n.Insert(11)
	n.Insert(55)
	n.Insert(77)

	n.Inorder()
	fmt.Println()

	n.Delete(11)

	n.Inorder()
	fmt.Println()

}
