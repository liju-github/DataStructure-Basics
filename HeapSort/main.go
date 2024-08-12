package main

import (
	"fmt"
)

func heapify(arr []int, heapSize int, rootIndex int) {
	largest := rootIndex
	leftChild := 2*rootIndex + 1
	rightChild := 2*rootIndex + 2

	if leftChild < heapSize && arr[leftChild] > arr[largest] {
		largest = leftChild
	}

	if rightChild < heapSize && arr[rightChild] > arr[largest] {
		largest = rightChild
	}

	if largest != rootIndex {
		arr[rootIndex], arr[largest] = arr[largest], arr[rootIndex]
		heapify(arr, heapSize, largest)
	}
}

func heapSort(arr []int) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Original array:", arr)

	heapSort(arr)
	fmt.Println("Sorted array:", arr)
}
