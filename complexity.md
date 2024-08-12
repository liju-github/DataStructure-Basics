Here's a breakdown of the time and space complexities for different data structures and algorithms:

### **1. Binary Tree**

- **Time Complexity:**
  - Insertion: `O(h)` (where `h` is the height of the tree)
  - Deletion: `O(h)`
  - Search: `O(h)`
  
  In a balanced binary tree, `h = O(log n)`, but in the worst case (skewed tree), `h = O(n)`.

- **Space Complexity:**
  - `O(n)` (for storing `n` nodes)

### **2. Binary Search Tree (BST)**

- **Time Complexity:**
  - Insertion: `O(h)`
  - Deletion: `O(h)`
  - Search: `O(h)`
  
  In a balanced BST, `h = O(log n)`, but in the worst case (unbalanced BST), `h = O(n)`.

- **Space Complexity:**
  - `O(n)` (for storing `n` nodes)

### **3. Heap (Min and Max)**

- **Time Complexity:**
  - Insertion: `O(log n)`
  - Deletion (removal of root): `O(log n)`
  - Accessing Min/Max: `O(1)`
  - Building a heap (Heapify): `O(n)`

- **Space Complexity:**
  - `O(n)` (for storing `n` elements)

### **4. Heap Sort**

- **Time Complexity:**
  - Building a heap: `O(n)`
  - Sorting: `O(n log n)`
  
  Overall: `O(n log n)`

- **Space Complexity:**
  - `O(1)` (in-place sorting, no additional space needed besides the input array)

### **5. Graph**

- **Time Complexity:**
  - Depending on the representation:
    - **Adjacency Matrix:**
      - Adding an edge: `O(1)`
      - Removing an edge: `O(1)`
      - Checking for an edge: `O(1)`
      - Space: `O(V^2)` (where `V` is the number of vertices)
    - **Adjacency List:**
      - Adding an edge: `O(1)`
      - Removing an edge: `O(E)` (where `E` is the number of edges)
      - Checking for an edge: `O(V)` (in the worst case)
      - Space: `O(V + E)`

- **Space Complexity:**
  - **Adjacency Matrix:** `O(V^2)`
  - **Adjacency List:** `O(V + E)`

### **6. Graph Traversals (BFS and DFS)**

- **Time Complexity:**
  - **BFS (Breadth-First Search):** `O(V + E)`
  - **DFS (Depth-First Search):** `O(V + E)`

- **Space Complexity:**
  - **BFS:** `O(V)` (for the queue and visited array)
  - **DFS:** `O(V)` (for the stack and visited array)

### **7. Trie**

- **Time Complexity:**
  - Insertion: `O(m)` (where `m` is the length of the key)
  - Deletion: `O(m)`
  - Search: `O(m)`

- **Space Complexity:**
  - `O(26 * m * n)` (for storing `n` keys with a maximum length of `m`, assuming each node has 26 possible children for the alphabet)

### **Summary:**

- **Binary Tree:**
  - **Time Complexity:** `O(h)` for insertion, deletion, and search.
  - **Space Complexity:** `O(n)`
  
- **Binary Search Tree:**
  - **Time Complexity:** `O(h)` for insertion, deletion, and search.
  - **Space Complexity:** `O(n)`
  
- **Heap (Min and Max):**
  - **Time Complexity:** `O(log n)` for insertion and deletion, `O(1)` for access.
  - **Space Complexity:** `O(n)`
  
- **Heap Sort:**
  - **Time Complexity:** `O(n log n)`
  - **Space Complexity:** `O(1)`
  
- **Graph:**
  - **Time Complexity:** Varies depending on representation and operations.
  - **Space Complexity:** `O(V^2)` for adjacency matrix, `O(V + E)` for adjacency list.
  
- **Graph Traversals (BFS and DFS):**
  - **Time Complexity:** `O(V + E)`
  - **Space Complexity:** `O(V)`
  
- **Trie:**
  - **Time Complexity:** `O(m)` for insertion, deletion, and search.
  - **Space Complexity:** `O(26 * m * n)`