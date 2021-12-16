package main

import "fmt"

func binarySearch(arr []int, x int) (bool, int) {
	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := int((start + end) / 2)
		if arr[mid] == x {
			return true, mid
		} else if x > arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false, -1
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
