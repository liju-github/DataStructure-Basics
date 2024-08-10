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


func (n *Node)Height() int  {
	if n == nil{
		return -1
	}

	LeftHeight := n.Left.Height()
	RightHeight := n.Right.Height()

	return max(LeftHeight,RightHeight)+1
}

func main()  {
	root := &Node{Value: 100}

	root.Insert(90)
	root.Insert(80)
	root.Insert(70)
	root.Insert(60)
	root.Insert(50)
	root.Insert(40)
	root.Insert(30)

	fmt.Println("Height of the tree is: ",root.Height())
}