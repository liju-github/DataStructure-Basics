package main

import (
	"fmt"
)

type MaxHeap struct {
	array []int
	size  int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		array: make([]int, 0),
		size:  0,
	}
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) leftChild(i int) int {
	return 2*i + 1
}

func (h *MaxHeap) rightChild(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.size++
	h.heapifyUp(h.size - 1)
}

func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 && h.array[h.parent(index)] < h.array[index] {
		h.swap(index, h.parent(index))
		index = h.parent(index)
	}
}

func (h *MaxHeap) Remove() (int, bool) {
	if h.size == 0 {
		return 0, false
	}

	max := h.array[0]
	h.array[0] = h.array[h.size-1]
	h.array = h.array[:h.size-1]
	h.size--

	if h.size > 0 {
		h.heapifyDown(0)
	}

	return max, true
}

func (h *MaxHeap) heapifyDown(index int) {
	largest := index
	left := h.leftChild(index)
	right := h.rightChild(index)

	if left < h.size && h.array[left] > h.array[largest] {
		largest = left
	}

	if right < h.size && h.array[right] > h.array[largest] {
		largest = right
	}

	if largest != index {
		h.swap(index, largest)
		h.heapifyDown(largest)
	}
}

func (h *MaxHeap) BuildHeap(arr []int) {
	h.array = arr
	h.size = len(arr)
	for i := h.size/2 - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

func main() {
	fmt.Println("Max Heap Example:")
	maxHeap := NewMaxHeap()
	maxHeap.Insert(3)
	maxHeap.Insert(2)
	maxHeap.Insert(1)
	maxHeap.Insert(5)
	maxHeap.Insert(4)

	fmt.Println("Removing elements from Max Heap:")
	for maxHeap.size > 0 {
		max, _ := maxHeap.Remove()
		fmt.Printf("%d ", max)
	}
	fmt.Println()

	fmt.Println("\nBuilding Max Heap from array:")
	arr := []int{10, 5, 15, 2, 20, 1, 30, 3, 4, 25}
	maxHeap.BuildHeap(arr)
	fmt.Println("Removing elements from built Max Heap:")
	for maxHeap.size > 0 {
		max, _ := maxHeap.Remove()
		fmt.Printf("%d ", max)
	}
	fmt.Println()
}