package main

import (
	"fmt"
)

// MinHeap represents a Min Heap data structure
// where the smallest element is always at the root.
type MinHeap struct {
	array []int // Array to store heap elements
	size  int   // Current size of the heap
}

// NewMinHeap creates a new MinHeap instance.
func NewMinHeap() *MinHeap {
	return &MinHeap{
		array: make([]int, 0),
		size:  0,
	}
}

// parent returns the index of the parent of the given node.
func (h *MinHeap) parent(i int) int {
	return (i - 1) / 2
}

// leftChild returns the index of the left child of the given node.
func (h *MinHeap) leftChild(i int) int {
	return 2*i + 1
}

// rightChild returns the index of the right child of the given node.
func (h *MinHeap) rightChild(i int) int {
	return 2*i + 2
}

// swap exchanges the values at two indices in the heap array.
func (h *MinHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

// Insert adds a new key to the Min Heap and maintains the heap property.
func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key) // Add the new key to the end of the array
	h.size++                       // Increase heap size
	h.heapifyUp(h.size - 1)        // Restore heap property from the bottom up
}

// heapifyUp moves the element at the given index up to restore the heap property.
func (h *MinHeap) heapifyUp(index int) {
	for index > 0 && h.array[h.parent(index)] > h.array[index] {
		h.swap(index, h.parent(index)) // Swap with parent if smaller
		index = h.parent(index)        // Move up to the parent index
	}
}

// Remove removes and returns the smallest element (root) of the Min Heap.
// Returns false if the heap is empty.
func (h *MinHeap) Remove() (int, bool) {
	if h.size == 0 {
		return 0, false // Heap is empty
	}

	min := h.array[0]              // Store the smallest element
	h.array[0] = h.array[h.size-1] // Replace root with the last element
	h.array = h.array[:h.size-1]   // Shrink the array
	h.size--                       // Decrease heap size

	if h.size > 0 {
		h.heapifyDown(0) // Restore heap property from the top down
	}

	return min, true // Return the smallest element
}

// heapifyDown moves the element at the given index down to restore the heap property.
func (h *MinHeap) heapifyDown(index int) {
	smallest := index            // Assume the smallest is at the current index
	left := h.leftChild(index)   // Index of the left child
	right := h.rightChild(index) // Index of the right child

	// Check if left child is smaller than current smallest
	if left < h.size && h.array[left] < h.array[smallest] {
		smallest = left
	}

	// Check if right child is smaller than current smallest
	if right < h.size && h.array[right] < h.array[smallest] {
		smallest = right
	}

	// If the smallest is not the current index, swap and continue
	if smallest != index {
		h.swap(index, smallest)
		h.heapifyDown(smallest)
	}
}

// BuildHeap creates a Min Heap from an existing array.
func (h *MinHeap) BuildHeap(arr []int) {
	h.array = arr     // Use the input array as the heap's internal array
	h.size = len(arr) // Update the size of the heap

	// Start heapifying from the last non-leaf node to the root
	for i := h.size/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func main() {
	fmt.Println("Min Heap Example:")

	// Create a new Min Heap
	minHeap := NewMinHeap()

	// Insert elements into the heap
	minHeap.Insert(3)
	minHeap.Insert(2)
	minHeap.Insert(1)
	minHeap.Insert(5)
	minHeap.Insert(4)

	fmt.Println("Removing elements from Min Heap:")
	for minHeap.size > 0 {
		min, _ := minHeap.Remove()
		fmt.Printf("%d ", min)
	}
	fmt.Println()

	fmt.Println("\nBuilding Min Heap from array:")

	// Build a Min Heap from an existing array
	arr := []int{10, 5, 15, 2, 20, 1, 30, 3, 4, 25}
	minHeap.BuildHeap(arr)

	fmt.Println("Removing elements from built Min Heap:")
	for minHeap.size > 0 {
		min, _ := minHeap.Remove()
		fmt.Printf("%d ", min)
	}
	fmt.Println()
}
