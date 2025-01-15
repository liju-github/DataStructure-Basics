


### **Binary Tree Operations Checklist**

1. **Insertion (Level-Order)**
   - **Condition**: Insert nodes in the first available spot in a level-order manner.
   - **Algorithm**:
     1. If the root is `nil`, insert the node as the root.
     2. Use a queue to traverse nodes level by level.
     3. Dequeue a node and check its left and right children.
     4. If a child is `nil`, insert the new node there.
     5. If not, enqueue the child and continue.

   ```go
   func (n *Node) Insert(value int) {
       newNode := &Node{Value: value}
       if n == nil {
           n = newNode
           return
       }
       queue := []*Node{n}
       for len(queue) > 0 {
           current := queue[0]
           queue = queue[1:]
           if current.Left == nil {
               current.Left = newNode
               return
           } else {
               queue = append(queue, current.Left)
           }
           if current.Right == nil {
               current.Right = newNode
               return
           } else {
               queue = append(queue, current.Right)
           }
       }
   }
   ```

2. **In-Order Traversal**
   - **Condition**: Left -> Node -> Right
   - **Algorithm**:
     1. Recursively traverse the left subtree.
     2. Visit the node.
     3. Recursively traverse the right subtree.

   ```go
   func (n *Node) Inorder() {
       if n == nil {
           return
       }
       n.Left.Inorder()
       fmt.Print(n.Value, " ")
       n.Right.Inorder()
   }
   ```

3. **Pre-Order Traversal**
   - **Condition**: Node -> Left -> Right
   - **Algorithm**:
     1. Visit the node.
     2. Recursively traverse the left subtree.
     3. Recursively traverse the right subtree.

   ```go
   func (n *Node) Preorder() {
       if n == nil {
           return
       }
       fmt.Print(n.Value, " ")
       n.Left.Preorder()
       n.Right.Preorder()
   }
   ```

4. **Post-Order Traversal**
   - **Condition**: Left -> Right -> Node
   - **Algorithm**:
     1. Recursively traverse the left subtree.
     2. Recursively traverse the right subtree.
     3. Visit the node.

   ```go
   func (n *Node) Postorder() {
       if n == nil {
           return
       }
       n.Left.Postorder()
       n.Right.Postorder()
       fmt.Print(n.Value, " ")
   }
   ```

5. **Level-Order Traversal**
   - **Condition**: Traverse nodes level by level using a queue.
   - **Algorithm**:
     1. Start with the root in the queue.
     2. Dequeue a node, visit it, and enqueue its children.
     3. Repeat until the queue is empty.

   ```go
   func (n *Node) LevelOrder() {
       if n == nil {
           return
       }
       queue := []*Node{n}
       for len(queue) > 0 {
           current := queue[0]
           queue = queue[1:]
           fmt.Print(current.Value, " ")
           if current.Left != nil {
               queue = append(queue, current.Left)
           }
           if current.Right != nil {
               queue = append(queue, current.Right)
           }
       }
   }
   ```

6. **Height Calculation**
   - **Condition**: Height of a node = max height of its left or right subtree + 1.
   - **Algorithm**:
     1. Base case: If node is `nil`, return `-1`.
     2. Recursively calculate the height of left and right subtrees.
     3. Return the maximum of these heights + 1.

   ```go
   func (n *Node) Height() int {
       if n == nil {
           return -1
       }
       leftHeight := n.Left.Height()
       rightHeight := n.Right.Height()
       return max(leftHeight, rightHeight) + 1
   }
   ```

7. **Depth Calculation**
   - **Condition**: Depth of a node = number of edges from the root to the node.
   - **Algorithm**:
     1. Base case: If node is `nil`, return `-1`.
     2. If node is found, return `0`.
     3. Recursively calculate the depth by searching left and right subtrees.
     4. If found, return depth + 1; otherwise, return `-1`.

   ```go
   func (n *Node) FindDepth(target int) int {
       if n == nil {
           return -1
       }
       if n.Value == target {
           return 0
       }
       leftDepth := n.Left.FindDepth(target)
       if leftDepth != -1 {
           return leftDepth + 1
       }
       rightDepth := n.Right.FindDepth(target)
       if rightDepth != -1 {
           return rightDepth + 1
       }
       return -1
   }
   ```

8. **Deletion (Any Node)**
   - **Condition**: Remove the node and fill its position with the deepest and rightmost node.
   - **Algorithm**:
     1. Perform a level-order traversal to find the deepest and rightmost node.
     2. Replace the node to be deleted with this node.
     3. Delete the deepest node.

   ```go
   func (n *Node) Delete(value int) {
       if n == nil {
           return
       }
       queue := []*Node{n}
       var targetNode, deepestNode *Node
       var parentOfDeepest *Node

       for len(queue) > 0 {
           current := queue[0]
           queue = queue[1:]

           if current.Value == value {
               targetNode = current
           }
           if current.Left != nil {
               parentOfDeepest = current
               queue = append(queue, current.Left)
           }
           if current.Right != nil {
               parentOfDeepest = current
               queue = append(queue, current.Right)
           }
           deepestNode = current
       }

       if targetNode != nil {
           targetNode.Value = deepestNode.Value
           if parentOfDeepest.Right == deepestNode {
               parentOfDeepest.Right = nil
           } else {
               parentOfDeepest.Left = nil
           }
       }
   }
   ```

---

### **Binary Search Tree (BST) Operations Checklist**

1. **Insertion**
   - **Condition**: Insert based on value comparison.
   - **Algorithm**:
     1. Start at the root.
     2. If the value is less than the current node, move to the left child.
     3. If the value is greater, move to the right child.
     4. Insert when a `nil` child is found.

   ```go
   func (n *Node) Insert(value int) {
       if n == nil {
           return
       }
       if value < n.Value {
           if n.Left == nil {
               n.Left = &Node{Value: value}
           } else {
               n.Left.Insert(value)
           }
       } else {
           if n.Right == nil {
               n.Right = &Node{Value: value}
           } else {
               n.Right.Insert(value)
           }
       }
   }
   ```

2. **Search**
   - **Condition**: Search based on value comparison.
   - **Algorithm**:
     1. Start at the root.
     2. If the value equals the current node, return true.
     3. If the value is less, search in the left subtree.
     4. If the value is greater, search in the right subtree.
     5. If the node is `nil`, return false.

   ```go
   func (n *Node) Search(value int) bool {
       if n == nil {
           return false
       }
       if value == n.Value {
           return true
       } else if value < n.Value {
           return n.Left.Search(value)
       } else {
           return n.Right.Search(value)
       }
   }
   ```

3. **Deletion**
   - **Condition**: Handle three cases: leaf node, one child, two children.
   - **Algorithm**:
     1. Find the node to be deleted.
     2. **Leaf Node**: Remove the node.
     3. **One Child**: Replace the node with its child.
     4. **Two Children**: Replace with in-order predecessor/successor, then delete it.

   ```go
   func (n *Node) Delete(value int) *Node {
       if n == nil {
           return nil
       }

       if value < n.Value {
           n.Left = n.Left.Delete(value)
       } else if value > n.Value {
           n.Right = n.Right.Delete(value)
       } else {
           if n.Left == nil {
               return n.Right
           } else if n.Right == nil {
               return n.Left
           }

           minRight := n.Right.minValue()
           n.Value = minRight.Value
           n.Right = n.Right.Delete(minRight.Value)
       }
       return n
   }

   func (n *Node) minValue() *Node {
       current:= n
       for current.Left != nil {
           current = current.Left
       }
       return current
   }
   ```

4. **In-Order Traversal**
   - **Condition**: Traverse Left -> Node -> Right.
   - **Algorithm**: Same as Binary Tree.

   ```go
   // Use the same Inorder function as in the Binary Tree
   ```

5. **Height Calculation**
   - **Condition**: Same as in Binary Tree.
   - **Algorithm**: Same as Binary Tree.

   ```go
   // Use the same Height function as in the Binary Tree
   ```

6. **Depth Calculation**
   - **Condition**: Same as Binary Tree but follows BST properties.
   - **Algorithm**: Use comparison to decide the direction (left or right).

   ```go
   func (n *Node) FindDepth(target int) int {
       if n == nil {
           return -1
       }
       if n.Value == target {
           return 0
       }
       if target < n.Value {
           leftDepth := n.Left.FindDepth(target)
           if leftDepth != -1 {
               return leftDepth + 1
           }
       } else {
           rightDepth := n.Right.FindDepth(target)
           if rightDepth != -1 {
               return rightDepth + 1
           }
       }
       return -1
   }
   ```

7. **Validate BST**
   - **Condition**: Ensure the left subtree values are less than the node, and the right subtree values are greater.
   - **Algorithm**:
     1. Recursively check the validity of left and right subtrees.
     2. Compare node values with allowable min and max values.

   ```go
   func (n *Node) Validate(min, max int) bool {
       if n == nil {
           return true
       }
       if n.Value <= min || n.Value >= max {
           return false
       }
       return n.Left.Validate(min, n.Value) && n.Right.Validate(n.Value, max)
   }
   ```

8. **Find Closest Value**
   - **Condition**: Traverse towards the closest value by comparing node values.
   - **Algorithm**:
     1. Start at the root and keep track of the closest value.
     2. Move left or right based on comparison with the target value.
     3. Update the closest value as needed.

   ```go
   func (n *Node) FindClosest(target int) int {
       closest := n.Value
       currentNode := n
       for currentNode != nil {
           if abs(target-currentNode.Value) < abs(target-closest) {
               closest = currentNode.Value
           }
           if target < currentNode.Value {
               currentNode = currentNode.Left
           } else if target > currentNode.Value {
               currentNode = currentNode.Right
           } else {
               break
           }
       }
       return closest
   }
   ```

---

### **Additional Notes**

- **Traversal**: In a Binary Tree, traversal methods (In-order, Pre-order, Post-order) don’t consider node values. In a BST, In-order traversal results in a sorted list of values.
- **Complexity**:
  - Binary Tree operations often have a time complexity of O(n) due to the lack of structure.
  - BST operations can have a time complexity of O(log n) if balanced, but O(n) in the worst case (unbalanced).
