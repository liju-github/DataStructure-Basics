package main

import "fmt"

type Node struct{
	Value int
	Left *Node
	Right *Node
}

func (n *Node)Insert(value int)  {
	newnode := &Node{Value: value}

	if value < n.Value{
		if n.Left == nil{
			n.Left = newnode
		}else{
			n.Left.Insert(value)
		}
	}else{
		if n.Right == nil{
			n.Right = newnode
		}else{
			n.Right.Insert(value)
		}
	}
}


func (n *Node)Inorder()  {
	if n ==nil{
		return

	}

	n.Left.Inorder()
	fmt.Print("->",n.Value)
	n.Right.Inorder()
}

func (n *Node)Height() int {
	if n == nil{
        return -1
	}

	leftHeight := n.Left.Height()
	rightHeight := n.Right.Height()

	return max(leftHeight,rightHeight)+1
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

	fmt.Println("Height is ",n.Height())

}