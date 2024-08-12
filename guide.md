# Tree and Graph Data Structures Cheat Sheet

## 1. Trees

### Basic Concepts
- A tree is a hierarchical data structure consisting of nodes connected by edges.
- The topmost node is called the root.
- Nodes with no children are called leaves.
- The depth of a node is its distance from the root.
- The height of a tree is the length of the path from the root to the deepest leaf.

### Types of Trees
1. Binary Tree: Each node has at most two children.
2. Binary Search Tree (BST): A binary tree where for each node, all elements in the left subtree are smaller, and all elements in the right subtree are larger.
3. AVL Tree: A self-balancing BST.
4. Red-Black Tree: Another type of self-balancing BST.

### Tree Traversals
1. In-order: Left subtree, Root, Right subtree
2. Pre-order: Root, Left subtree, Right subtree
3. Post-order: Left subtree, Right subtree, Root
4. Level-order: Visit nodes level by level

## 2. Binary Search Tree (BST)

### Properties
- For each node, all elements in the left subtree are smaller, and all elements in the right subtree are larger.
- Allows for efficient searching, insertion, and deletion operations.

### Operations
1. Insertion: O(log n) average case, O(n) worst case
2. Deletion: O(log n) average case, O(n) worst case
3. Search: O(log n) average case, O(n) worst case

### BST Validation
To validate if a binary tree is a BST:
1. Keep track of the allowed range for each node.
2. Recursively check if each node's value is within its allowed range.
3. Update the range for left and right children accordingly.

### Finding Closest Value
To find the closest value to a given number:
1. Start at the root and initialize the closest value.
2. Traverse the tree, updating the closest value if a closer one is found.
3. Move left if the target is smaller than the current node, right if larger.

## 3. Heaps

### Basic Concepts
- A heap is a complete binary tree where each node satisfies the heap property.
- Max Heap: Parent nodes are always greater than or equal to their children.
- Min Heap: Parent nodes are always smaller than or equal to their children.

### Operations
1. Build Heap: O(n)
2. Insert: O(log n)
3. Remove (Extract Min/Max): O(log n)
4. Peek (Get Min/Max): O(1)

### Heap Implementation
- Typically implemented using an array.
- For a node at index i:
  - Left child: 2i + 1
  - Right child: 2i + 2
  - Parent: (i - 1) / 2

## 4. Heap Sort

### Concept
- Uses a max heap to sort elements in ascending order.
- Time Complexity: O(n log n) for all cases.
- Space Complexity: O(1) as it sorts in-place.

### Steps
1. Build a max heap from the input array.
2. Swap the root (largest element) with the last element of the heap.
3. Reduce heap size by 1 and heapify the root.
4. Repeat steps 2-3 until the heap size becomes 1.

## 5. Trie (Prefix Tree)

### Basic Concepts
- A tree-like data structure used to store and retrieve strings.
- Each node represents a character in a string.
- The root represents an empty string.
- Each path from root to a node forms a prefix of one or more strings.

### Operations
1. Insert: O(m), where m is the length of the string
2. Search: O(m)
3. Delete: O(m)
4. Prefix Search: O(p), where p is the length of the prefix

### Applications
- Autocomplete
- Spell checkers
- IP routing tables
- Longest prefix matching

## 6. Graphs

### Basic Concepts
- A graph consists of vertices (nodes) and edges connecting these vertices.
- Directed Graph: Edges have a direction.
- Undirected Graph: Edges have no direction.
- Weighted Graph: Edges have associated weights or costs.

### Representations
1. Adjacency Matrix: 2D array where A[i][j] represents an edge from vertex i to j.
2. Adjacency List: Array of lists, where each list contains the neighbors of a vertex.

### Graph Traversals

#### Breadth-First Search (BFS)
- Explores all vertices at the present depth before moving to vertices at the next depth level.
- Uses a queue data structure.
- Time Complexity: O(V + E), where V is the number of vertices and E is the number of edges.

#### Depth-First Search (DFS)
- Explores as far as possible along each branch before backtracking.
- Uses a stack or recursion.
- Time Complexity: O(V + E)

### Applications
- Shortest path algorithms (e.g., Dijkstra's algorithm)
- Minimum spanning tree (e.g., Kruskal's algorithm, Prim's algorithm)
- Cycle detection
- Topological sorting

This cheat sheet covers the fundamental concepts of trees, binary search trees, heaps, heap sort, tries, and graphs, including their basic operations and time complexities. It provides a solid theoretical foundation for understanding these data structures and their applications.