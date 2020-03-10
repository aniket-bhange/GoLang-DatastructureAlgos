package sort

import (
	"DS/utils"
)

//Heapify to find largest element
func Heapify(arr []int, n int, i int) []int {
	largest := i
	lChild := 2*i + 1
	rChild := 2*i + 2

	if lChild < n && arr[lChild] > arr[largest] {
		largest = lChild
	}

	if rChild < n && arr[rChild] > arr[largest] {
		largest = rChild
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]

		arr = Heapify(arr, n, largest)
	}
	return arr
}

//HeapSort function to make heap and hepify elements
func HeapSort(arr []int) []int {
	length := len(arr)

	for i := length/2 - 1; i >= 0; i-- {
		arr = Heapify(arr, length, i)
	}

	for i := length - 1; i >= 0; i-- {
		utils.Swap(&arr[0], &arr[i])
		arr = Heapify(arr, i, 0)
	}
	return arr
}
