package main

import (
	"fmt"
)

type MinHeap struct {
	array []int
	size  int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		array: make([]int, 0),
		size:  0,
	}
}

func (h *MinHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) leftChild(i int) int {
	return 2*i + 1
}

func (h *MinHeap) rightChild(i int) int {
	return 2*i + 2
}

func (h *MinHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.size++
	h.heapifyUp(h.size - 1)
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 && h.array[h.parent(index)] > h.array[index] {
		h.swap(index, h.parent(index))
		index = h.parent(index)
	}
}

func (h *MinHeap) Remove() (int, bool) {
	if h.size == 0 {
		return 0, false
	}

	min := h.array[0]
	h.array[0] = h.array[h.size-1]
	h.array = h.array[:h.size-1]
	h.size--

	if h.size > 0 {
		h.heapifyDown(0)
	}

	return min, true
}

func (h *MinHeap) heapifyDown(index int) {
	smallest := index
	left := h.leftChild(index)
	right := h.rightChild(index)

	if left < h.size && h.array[left] < h.array[smallest] {
		smallest = left
	}

	if right < h.size && h.array[right] < h.array[smallest] {
		smallest = right
	}

	if smallest != index {
		h.swap(index, smallest)
		h.heapifyDown(smallest)
	}
}

func (h *MinHeap) BuildHeap(arr []int) {
	h.array = arr
	h.size = len(arr)
	for i := h.size/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func main() {
	fmt.Println("Min Heap Example:")
	minHeap := NewMinHeap()
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
	arr := []int{10, 5, 15, 2, 20, 1, 30, 3, 4, 25}
	minHeap.BuildHeap(arr)
	fmt.Println("Removing elements from built Min Heap:")
	for minHeap.size > 0 {
		min, _ := minHeap.Remove()
		fmt.Printf("%d ", min)
	}
	fmt.Println()
}