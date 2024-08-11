package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value int) {
	newnode := &Node{Value: value}

	if n == nil{
		n = newnode
	}

	queue := []*Node{n}

	for len(queue) > 0{
		current := queue[0]
		queue = queue[1:]

		if current.Left == nil{
			current.Left = newnode
			break
		}else{
			queue = append(queue, current.Left)
		}

		if current.Right == nil{
			current.Right = newnode
			break
		}else{
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

func (n *Node) Height() int {
	if n == nil {
		return -1
	}

	leftHeight := n.Left.Height() 
	rightHeight := n.Right.Height() 

	return max(leftHeight, rightHeight) + 1

}

func main() {
	n := &Node{Value: 100}

	n.Insert(90)
	n.Insert(85)
	n.Insert(70)
	n.Insert(110)
	n.Insert(80)
	n.Insert(40)
	n.Insert(220)
	n.Insert(50)
	n.Insert(20)

	n.Inorder()
	fmt.Println()
	fmt.Println("Height of the node : ",n.Height())
}
