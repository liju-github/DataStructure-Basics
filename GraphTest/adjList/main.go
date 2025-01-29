package main


import "fmt"

type Graph struct{
	adjList map[int][]int
}

func NewGraph()*Graph{
	return &Graph{
		adjList: make(map[int][]int),
	}
}

func main(){
	g := NewGraph()
	g.Insert(1,2)
	g.Insert(2,3)
	g.Insert(3,4)
	g.Insert(6,7)
	g.Insert(7,8)

	g.DFS()
	g.BFS()
}

func (g *Graph)Insert(a, b int){
	g.adjList[a] = append(g.adjList[a],b)
	g.adjList[b] = append(g.adjList[b],a)
}

func (g *Graph)DFS(){
	visited:= make(map[int]bool)

	var dfsHelper func (start int)
	dfsHelper = func (start int)  {
		visited[start] = true
		fmt.Print(" -> ",start)
		for _,neigh := range g.adjList[start]{
	    if !visited[neigh]{
				dfsHelper(neigh)
			}
		}
	}

	for v := range g.adjList{
		fmt.Print("\nComponent : ",v)
		dfsHelper(v)
	}
}


func (g *Graph)BFS(){
	visited := make(map[int]bool)
	fmt.Println("\nbfs")


	for  v := range g.adjList{
		if !visited[v]{
			fmt.Print(" Component : ",v)
			g.BFShelper(v,visited)
		}
	}
}

func (g *Graph)BFShelper(node int,visited map[int]bool){
	queue := []int{node}
	visited[node] = true
	for len(queue) > 0{
		curr := queue[0]
		queue = queue[1:]

		fmt.Print(" -> ",curr)

		for _,v := range g.adjList[curr]{
			if !visited[v]{
				visited[v] = true
			queue = append(queue, v)
			}
		}
	}

}