package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value int) {

	newnode := &Node{Value: value}

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
    if n == nil{
		return
	}

	n.Left.Inorder()
	fmt.Print("-> ",n.Value)
    n.Right.Inorder()
}

func (n *Node) PreOrder() {
   if n == nil{
	return
   }

   fmt.Print("-> ",n.Value)
   n.Left.PreOrder()
   n.Right.PreOrder()
}

func (n *Node) PostOrder() {
    if n == nil{
		return
	}

	n.Left.PostOrder()
	n.Right.PostOrder()
	fmt.Print("-> ",n.Value)
}


func main()  {
	n := &Node{Value: 100}

	n.Insert(90)
	n.Insert(80)
	n.Insert(70)
	n.Insert(60)
	n.Insert(50)
	n.Insert(40)
	n.Insert(30)
	n.Insert(20)
	n.Insert(10)

	fmt.Println("Inorder Traversal : ")
	n.Inorder()
    println()

	fmt.Println("Preorder Traversal : ")
	n.PreOrder()
    println()

	fmt.Println("Postorder Traversal : ")
	n.PostOrder()
    println()
}