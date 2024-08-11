package main

import "fmt"


type Node struct{
	Value int
	Left *Node
	Right *Node
}

func (n *Node)Insert(value int)  {
	newnode := &Node{Value: value}

	if n == nil{
		n = newnode
	}

	queue := []*Node{n}

	for len(queue)>0{
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

func (n *Node)Inorder()  {
	if n==nil{
		return
	}

	n.Left.Inorder()
	fmt.Print("->",n.Value)
	n.Right.Inorder()
}

func (n *Node)Depth(target int) int  {
	if n == nil{
		return -1
	}

	if n.Value == target{
		return 0
	}

	LeftDepth := n.Left.Depth(target)
	if LeftDepth != -1 {
		return LeftDepth+1
	}

	RightDepth := n.Right.Depth(target)
	if RightDepth != -1{
		return RightDepth+1
	}

	return -1
}

func main()  {
	n := &Node{Value: 100}

	n.Insert(34)
	n.Insert(22)
	n.Insert(91)
	n.Insert(19)
	n.Insert(73)
	n.Insert(62)
	n.Insert(90)
	n.Insert(32)
	n.Insert(72)

	n.Inorder()
	fmt.Println()
	target := 90
	fmt.Println("depth of ",target,": ",n.Depth(target))
}