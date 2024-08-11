package main

import "fmt"


type Node struct{
	Value int
	Left *Node
	Right *Node
}


func (n *Node)Insert(value int)  {
	newnode:= &Node{Value: value}

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


func (n *Node)InorderTraversal()  {
	if n == nil{
		return
	}
	n.Left.InorderTraversal()
	fmt.Println("--> ",n.Value)
	n.Right.InorderTraversal()
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

	root.InorderTraversal()
}