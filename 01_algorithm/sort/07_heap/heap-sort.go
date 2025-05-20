package main

import "fmt"

func heapify(arr []int, n int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
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

	fmt.Println("----------------------------------------")
	arr0 := []int{1}
	fmt.Println("Original array:", arr0)
	heapSort(arr0)
	fmt.Println("Sorted array:", arr0)
	fmt.Println()

	fmt.Println("----------------------------------------")
	arr1 := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	fmt.Println("Original array:", arr1)
	heapSort(arr1)
	fmt.Println("Sorted array:", arr1)
}
