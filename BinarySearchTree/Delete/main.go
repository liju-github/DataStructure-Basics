package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Delete(value int) *Node {
	if n == nil {
		return nil
	}

	if value < n.Value {
		n.Left = n.Left.Delete(value)
	} else if value > n.Value {
		n.Right = n.Right.Delete(value)
	} else {
		if n.Left == nil && n.Right == nil {
			return nil
		}

		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		smallestRight := n.Right.findMin()
		n.Value = smallestRight.Value
		n.Right = n.Right.Delete(smallestRight.Value)
	}

	return n
}

func (n *Node) findMin() *Node {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (n *Node) Insert(value int) {
	newnode := &Node{Value: value}

	if n.Value > value {
		if n.Left == nil {
			n.Left = newnode
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = newnode
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *Node)Inorder()  {
	if n == nil{
		return
	}

	n.Left.Inorder()
	fmt.Print("->",n.Value)
	n.Right.Inorder()
}


func main()  {
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
