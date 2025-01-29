package main

/*
import "fmt"

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(a, b int) {
	g.adjList[a] = append(g.adjList[a], b)
	g.adjList[b] = append(g.adjList[b], a)
}

// DFS to visit all components
func (g *Graph) DFS() {
	visited := make(map[int]bool)
	fmt.Println("DFS Traversal:")
	for v := range g.adjList {
		if !visited[v] {
			fmt.Print("\nComponent: ")
			g.DFSRecursive(visited, v)
		}
	}
	fmt.Println()
}

func (g *Graph) DFSRecursive(visited map[int]bool, node int) {
	if visited[node] {
		return
	}

	fmt.Printf("-> %v ", node)
	visited[node] = true

	for _, v := range g.adjList[node] {
		g.DFSRecursive(visited, v)
	}
}

// BFS to visit all components
func (g *Graph) BFS() {
	visited := make(map[int]bool)
	fmt.Println("BFS Traversal:")
	for v := range g.adjList {
		if !visited[v] {
			fmt.Print("\nComponent: ")
			g.BFSHelper(v, visited)
		}
	}
	fmt.Println()
}

func (g *Graph) BFSHelper(start int, visited map[int]bool) {
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		fmt.Printf("-> %v ", curr) // Print current node

		for _, v := range g.adjList[curr] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
}

func main() {
	g := NewGraph()

	// Creating a graph with multiple connected components
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(6, 1)

	g.DFS()
	g.BFS()

	arr:= g.ShortestPath(1,5)

	fmt.Println(arr)
}

func (g *Graph)ShortestPath(start,end int)[]int{
	visited := make(map[int]bool)
	parent := make(map[int]int)
	queue := []int{start}

	reconstructPath := func (parent map[int]int,start,end int)[]int {
		path := []int{}

		for node := end; node != start; node = parent[node]{
			path = append([]int{node},path...)
		}

		path = append([]int{start},path...)
		return path
	}

	for len(queue) > 0{
		curr := queue[0]
		queue = queue[1:]

		for _ ,neighbour :=  range g.adjList[curr]{
			if !visited[neighbour]{
				visited[neighbour] = true
				parent[neighbour] = curr
				queue = append(queue, neighbour)
			}

			if neighbour == end{
				return reconstructPath(parent,start,end)
			}
		}
		
	}
	return nil
}
	*/