package main

import "fmt"

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return index*2 + 1
}

func right(index int) int {
	return index*2 + 2
}

func heapify_min(array *[]int) {
	arr := *array
	// first start rightIndex is the most right index of an array
	rightIndex := len(arr) - 1
	// current parent
	p := parent(rightIndex)

	for arr[p] > arr[rightIndex] {
		arr[p], arr[rightIndex] = arr[rightIndex], arr[p]
		// new rightIndex is current parent
		rightIndex = p
		// then p become next parent
		p = parent(rightIndex)
	}
}

func main() {
	array := []int{4}
	heapify_min(&array)
	fmt.Println(array)

	array = append(array, 3)
	heapify_min(&array)
	fmt.Println(array)

	array = append(array, 2)
	heapify_min(&array)
	fmt.Println(array)

	array = append(array, 3)
	heapify_min(&array)
	fmt.Println(array)

	array = append(array, 1)
	heapify_min(&array)
	fmt.Println(array)

	fmt.Println(array[left(1)])
	fmt.Println(array[right(1)])
}
