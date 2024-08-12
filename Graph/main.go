package main

import (
	"fmt"
)

type Graph struct {
	vertices map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

func (g *Graph) AddVertex(v int) {
	if _, exists := g.vertices[v]; !exists {
		g.vertices[v] = []int{}
	}
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.vertices[v1] = append(g.vertices[v1], v2)
	g.vertices[v2] = append(g.vertices[v2], v1)
}

func (g *Graph) RemoveEdge(v1, v2 int) {
	g.vertices[v1] = removeFromSlice(g.vertices[v1], v2)
	g.vertices[v2] = removeFromSlice(g.vertices[v2], v1)
}

func (g *Graph) RemoveVertex(v int) {
	delete(g.vertices, v)
	for vertex, edges := range g.vertices {
		g.vertices[vertex] = removeFromSlice(edges, v)
	}
}

func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	var result []int
	g.dfsUtil(start, visited, &result)
	return result
}

func (g *Graph) dfsUtil(v int, visited map[int]bool, result *[]int) {
	visited[v] = true
	*result = append(*result, v)
	for _, neighbor := range g.vertices[v] {
		if !visited[neighbor] {
			g.dfsUtil(neighbor, visited, result)
		}
	}
}

func (g *Graph) BFS(start int) []int {
	visited := make(map[int]bool)
	var result []int
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		result = append(result, v)

		for _, neighbor := range g.vertices[v] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

func removeFromSlice(slice []int, element int) []int {
	for i, v := range slice {
		if v == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func main() {
	graph := NewGraph()

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)

	fmt.Println("Graph:", graph.vertices)

	fmt.Println("DFS starting from vertex 0:", graph.DFS(0))

	fmt.Println("BFS starting from vertex 0:", graph.BFS(0))

	graph.RemoveEdge(1, 2)
	fmt.Println("Graph after removing edge (1,2):", graph.vertices)

	graph.RemoveVertex(2)
	fmt.Println("Graph after removing vertex 2:", graph.vertices)
}
