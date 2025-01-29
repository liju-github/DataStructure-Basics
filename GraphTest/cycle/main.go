package main

import "fmt"

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[int][]int),
	}
}

func (g *Graph) Insert(a, b int) {
	g.adjList[a] = append(g.adjList[a], b)
	g.adjList[b] = append(g.adjList[b], a)
}

func (g *Graph) DetectCycle() bool {
	visited := make(map[int]bool)

	for node := range g.adjList {
		if !visited[node] {
			if g.dfsCycle(node, -1, visited) { // -1 means no parent for first node
				return true
			}
		}
	}
	return false
}

func (g *Graph) dfsCycle(node, parent int, visited map[int]bool) bool {
	visited[node] = true

	for _, neighbor := range g.adjList[node] {
		if !visited[neighbor] {
			if g.dfsCycle(neighbor, node, visited) {
				return true
			}
		} else if neighbor != parent {
			// If the neighbor is visited and NOT the parent, cycle detected
			return true
		}
	}
	return false
}

func main() {
	g := NewGraph()
	g.Insert(1, 2)
	g.Insert(2, 3)
	g.Insert(3, 4)
	g.Insert(4, 5)
	g.Insert(5, 2) // This creates a cycle

	if g.DetectCycle() {
		fmt.Println("Cycle detected in the graph.")
	} else {
		fmt.Println("No cycle in the graph.")
	}
}
