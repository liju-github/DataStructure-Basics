// package main

// import "fmt"

// type Node struct{
// 	Value int
// 	Left *Node
// 	Right *Node
// }


// func (n *Node)Insert(value int)  {
// 	newnode := &Node{Value: value}

// 	if n == nil{
// 		return
// 	}

// 	if n.Value < value{
// 		if n.Left == nil{
// 			n.Left = newnode
// 		}else{
// 			n.Left.Insert(value)
// 		}	
// 	}else{
// 		if n.Right == nil{
// 			n.Right = newnode
// 		}else{
// 			n.Right.Insert(value)
// 		}
// 	}
// }

// func (n *Node)Inorder()  {
	
// 	if n == nil{
// 		return
// 	}

// 	n.Left.Inorder()
// 	fmt.Print("->",n.Value)
// 	n.Right.Inorder()
// }

// func main()  {
// 	root := &Node{Value: 100}

// 	root.Insert(90)
// 	root.Insert(80)
// 	root.Insert(70)
// 	root.Insert(60)
// 	root.Insert(50)
// 	root.Insert(40)
// 	root.Insert(30)

// 	root.Inorder()
// 	fmt.Println()
// }

package main 

import ("fmt")

// type Tree struct{
//     root *Node
// }

type Node struct{
    Left *Node
    value int
    Right *Node
}

func (t *Node)Insert(value int){
    newnode := &Node{value:value}
    
    if t == nil{
        return
    }
    
    if t.value > value{
        if t.Left == nil{
            t.Left = newnode
        }else{
            t.Left.Insert(value)
        }
    }else{
        if t.Right == nil{
            t.Right = newnode
        }else{
            t.Right.Insert(value)
        }
    }
}

func (t *Node)Inorder(){
    if t ==nil{
        return
    }
    t.Left.Inorder()
    fmt.Print("->",t.value)
    t.Right.Inorder()
}


func main(){
    t := &Node{value : 10};
    
    t.Insert(10)
    t.Insert(50)
    t.Insert(100)
    t.Insert(340)
    t.Insert(20)
    t.Insert(90)
}


