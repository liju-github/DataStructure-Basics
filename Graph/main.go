package main

import "fmt"

// GraphNode represents a single node in the graph
type GraphNode struct {
	Value     int
	Neighbors []*GraphNode
}

// Graph represents the graph structure
type Graph struct {
	Nodes map[int]*GraphNode
}

// NewGraph creates and returns a new graph
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]*GraphNode),
	}
}

// AddNode adds a new node to the graph
func (g *Graph) AddNode(value int) { //AddVertex
	if _, exists := g.Nodes[value]; !exists {
		g.Nodes[value] = &GraphNode{Value: value}
	}
}

// AddEdge adds an undirected edge between two nodes
func (g *Graph) AddEdge(value1, value2 int) {
	g.AddNode(value1)
	g.AddNode(value2)

	node1 := g.Nodes[value1]
	node2 := g.Nodes[value2]

	node1.Neighbors = append(node1.Neighbors, node2)
	node2.Neighbors = append(node2.Neighbors, node1)
}

// DFS performs a Depth-First Search starting from a given node
func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	var result []int
	g.dfsUtil(g.Nodes[start], visited, &result)
	return result
}

func (g *Graph) dfsUtil(node *GraphNode, visited map[int]bool, result *[]int) {
	if node == nil || visited[node.Value] {
		return
	}

	visited[node.Value] = true
	*result = append(*result, node.Value)

	for _, neighbor := range node.Neighbors {
		g.dfsUtil(neighbor, visited, result)
	}
}

// BFS performs a Breadth-First Search starting from a given node
func (g *Graph) BFS(start int) []int {
	startNode := g.Nodes[start]
	if startNode == nil {
		return nil
	}

	visited := make(map[int]bool)
	var result []int
	queue := []*GraphNode{startNode}
	visited[startNode.Value] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Value)

		for _, neighbor := range node.Neighbors {
			if !visited[neighbor.Value] {
				visited[neighbor.Value] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

func main() {
	graph := NewGraph()

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)

	fmt.Println("DFS starting from node 0:", graph.DFS(0))
	fmt.Println("BFS starting from node 0:", graph.BFS(0))
}
