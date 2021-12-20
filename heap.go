package main

import "fmt"

type MinHeap struct {
	arr []int
}

func Parent(index int) int {
	return (index - 1) / 2
}

func Left(index int) int {
	return index*2 + 1
}

func Right(index int) int {
	return index*2 + 2
}

func (h *MinHeap) Insert(value int) {
	h.arr = append(h.arr, value)
	h.HeapifyUp()
}

func (h *MinHeap) Extract() int {
	if len(h.arr) == 0 {
		return -1
	}

	// rule of extraction
	lastIndex := len(h.arr) - 1
	extracted := h.arr[0]
	h.arr[0] = h.arr[lastIndex]
	h.arr = h.arr[:lastIndex]
	h.HeapifyDown()

	return extracted
}

// heapify from bottom to top
func (h *MinHeap) HeapifyUp() {
	// Initial startRightIndex is the most right index of an array
	startRightIndex := len(h.arr) - 1
	// Initial current active parent
	p := Parent(startRightIndex)

	for h.arr[p] > h.arr[startRightIndex] {
		h.arr[p], h.arr[startRightIndex] = h.arr[startRightIndex], h.arr[p]
		// next startRightIndex is current active parent
		startRightIndex = p
		// then p become next active parent
		p = Parent(startRightIndex)
	}
}

// heapify from top to bottom
func (h *MinHeap) HeapifyDown() {
	lastIndex := len(h.arr) - 1
	parentIndex := 0
	leftChildIndex, rightChildIndex := Left(parentIndex), Right(parentIndex)
	childIndexToCompare := 0

	for leftChildIndex <= lastIndex {
		if leftChildIndex == lastIndex { // when left child is the only child (have no right child)
			childIndexToCompare = leftChildIndex
		} else if leftChildIndex < rightChildIndex { // when left child is smaller than right child
			childIndexToCompare = leftChildIndex
		} else { // when left child is larger than right child
			childIndexToCompare = rightChildIndex
		}

		if h.arr[parentIndex] > h.arr[childIndexToCompare] {
			h.arr[parentIndex], h.arr[childIndexToCompare] = h.arr[childIndexToCompare], h.arr[parentIndex]
			parentIndex = childIndexToCompare
			leftChildIndex, rightChildIndex = Left(parentIndex), Right(parentIndex)
		} else {
			return
		}
	}
}

func main() {
	h := &MinHeap{}

	// Insert / push some data into the queue
	h.Insert(4)
	h.Insert(3)
	h.Insert(2)
	h.Insert(3)
	h.Insert(1)

	fmt.Println("Insert")
	fmt.Println(h.arr)

	// Extract / pop some data from the queue
	fmt.Println("1st extract:", h.Extract())
	fmt.Println("2nd extract:", h.Extract())
	fmt.Println("3rd extract:", h.Extract())

	fmt.Println("Extract")
	fmt.Println(h.arr)
}
