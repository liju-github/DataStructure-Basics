package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value int) {
	newNode := &Node{Value: value}

	if value < n.Value {
		if n.Left == nil {
			n.Left = newNode
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = newNode
		} else {
			n.Right.Insert(value)
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

func (n *Node) Preorder() {
	if n == nil {
		return
	}

	fmt.Print("->", n.Value)
	n.Left.Preorder()
	n.Right.Preorder()
}

func (n *Node) Postorder() {
	if n == nil {
		return
	}

	n.Left.Postorder()
	n.Right.Postorder()
	fmt.Print("->", n.Value)
}

func (n *Node) Levelorder() {
	if n == nil {
		return
	}

	queue := []*Node{n}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Print("->", current.Value)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

func main() {

	f, err := os.Create("memprofile.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	pprof.WriteHeapProfile(f)

	n := &Node{Value: 50}
	n.Insert(30)
	n.Insert(70)
	n.Insert(20)
	n.Insert(40)
	n.Insert(60)
	n.Insert(80)

	fmt.Print("Inorder traversal: ")
	n.Inorder()
	fmt.Println()

	fmt.Print("Preorder traversal: ")
	n.Preorder()
	fmt.Println()

	fmt.Print("Postorder traversal: ")
	n.Postorder()
	fmt.Println()

	fmt.Print("Level-order traversal: ")
	n.Levelorder()
	fmt.Println()
}
