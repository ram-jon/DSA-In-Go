package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		for j > 0 && arr[j] < arr[j-1] {
			temp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = temp
			j--
		}
	}
}

func main() {
	arr := []int{254, 32, 64, 54, 39, 10, 86, 5, 319, 98}
	fmt.Println("Array before sorting : ", arr)
	insertionSort(arr)
	fmt.Println("Array after sorting: ", arr)
}
