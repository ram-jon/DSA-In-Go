package main

import "fmt"

func binarySearchHelper(arr []int, x int, start int, end int) (bool, int) {
	if start > end {
		return false, -1
	}
	mid := int((start + end) / 2)

	if arr[mid] == x {
		return true, mid
	} else if x > arr[mid] {
		return binarySearchHelper(arr, x, mid+1, end)
	} else {
		return binarySearchHelper(arr, x, start, mid-1)
	}
}

func binarySearch(arr []int, x int) (bool, int) {
	start := 0
	end := len(arr) - 1
	return binarySearchHelper(arr, x, start, end)
}

func main() {
	arr := []int{10, 32, 39, 54, 86, 98, 319}
	x := 86
	if found, index := binarySearch(arr, x); found {
		fmt.Printf("Element %v found at index %v \n", x, index)
	} else {
		fmt.Println("Element not found!!")
	}
}
