### Understanding Graphs for Beginners

#### What is a Graph?

A **graph** is a collection of nodes (also called vertices) connected by edges. Think of it like a network where each point (node) can be connected to one or more other points (nodes) by lines (edges). Graphs are used to represent various structures, such as social networks, maps, and networks of roads.

- **Vertices (or Nodes)**: These are the points in the graph. For example, in a social network, each person could be a vertex.
- **Edges**: These are the connections between the vertices. In a social network, an edge might represent a friendship between two people.

#### Types of Graphs

   <!-- ![Graph](https://ipython-books.github.io/pages/chapter14_graphgeo/images/graphs.png) -->

1. **Undirected Graph**: The edges have no direction, meaning the connection is bidirectional. If there is an edge between nodes A and B, you can travel from A to B or from B to A.

2. **Directed Graph**: The edges have a direction, meaning the connection is one-way. If there's an edge from A to B, you can only travel from A to B, not the other way around.
   

3. **Weighted Graph**: In addition to direction, edges have weights (values). These could represent distances, costs, or any other value associated with moving from one vertex to another.

#### Basic Graph Terminology

- **Path**: A sequence of vertices connected by edges.
- **Cycle**: A path where the first and last vertices are the same.
- **Connected Graph**: A graph where there is a path between every pair of vertices.
- **Disconnected Graph**: A graph where some vertices are not connected.

#### Graph Representation

Graphs can be represented in several ways:

1. **Adjacency Matrix**: A 2D array where the element at row i and column j represents the presence (and possibly the weight) of an edge between vertices i and j.
2. **Adjacency List**: A collection of lists. Each vertex has a list of vertices it is connected to.

#### Flow of Execution in a Graph

Let’s consider an undirected graph as in the code example you shared. Here’s how the flow of execution works:

1. **Graph Creation**: 
   - You start by creating a new graph using `NewGraph()`. 
   - Vertices and edges are then added using methods like `AddVertex()` and `AddEdge()`.

2. **Adding Vertices and Edges**:
   - When you add an edge between two vertices (say A and B), you are essentially saying, “A is connected to B” and “B is connected to A” (since it's undirected).
   - This is stored in the adjacency list, where each vertex has a list of vertices it's connected to.

3. **Traversal**:
   - **Depth-First Search (DFS)**: DFS starts at one vertex and explores as far as possible along each branch before backtracking. It uses a stack or recursion. Imagine exploring a maze by taking a path until you can’t go further, then backtracking and trying another path.
   - **Breadth-First Search (BFS)**: BFS starts at one vertex and explores all its neighbors before moving on to their neighbors. It uses a queue. Imagine exploring a maze level by level, moving outwards from your start point.

4. **Removing Edges and Vertices**:
   - When you remove an edge or vertex, the graph structure is updated to reflect this change. This is useful when the connections between nodes can change over time.

#### Example Walkthrough

Let's walk through the code with an example to see the flow:

```go
graph := NewGraph()
```
- This creates an empty graph.

```go
graph.AddEdge(0, 1)
graph.AddEdge(0, 2)
graph.AddEdge(1, 2)
graph.AddEdge(2, 3)
```
- You add edges between the vertices: 0-1, 0-2, 1-2, and 2-3.
- The adjacency list will look like this:
  ```
  0 -> [1, 2]
  1 -> [0, 2]
  2 -> [0, 1, 3]
  3 -> [2]
  ```

**DFS (starting from vertex 0)**:
- The DFS would visit the nodes in the order: 0 → 1 → 2 → 3 (the exact order may vary depending on implementation details).
  
**BFS (starting from vertex 0)**:
- The BFS would visit the nodes in the order: 0 → 1 → 2 → 3.

**Removing an Edge**:
- If you remove the edge between 1 and 2, the adjacency list would update:
  ```
  0 -> [1, 2]
  1 -> [0]
  2 -> [0, 3]
  3 -> [2]
  ```

**Removing a Vertex**:
- If you remove vertex 2, all connections to and from 2 are removed:
  ```
  0 -> [1]
  1 -> [0]
  3 -> []
  ```

### Key Concepts & Algorithms:

1. **Depth-First Search (DFS)**:
   - Explores as far as possible before backtracking.
   - Useful for finding connected components, topological sorting, etc.
   - Time Complexity: `O(V + E)`, where V is the number of vertices and E is the number of edges.

2. **Breadth-First Search (BFS)**:
   - Explores neighbors level by level.
   - Useful for finding the shortest path in an unweighted graph.
   - Time Complexity: `O(V + E)`.

3. **Graph Traversal**:
   - Moving through the graph in some order.
   - Can be done using DFS, BFS, or other algorithms.

By understanding these basic concepts and walking through examples, you'll get a solid foundation in how graphs work and how to manipulate them in your programs. Feel free to ask more questions as you explore further!