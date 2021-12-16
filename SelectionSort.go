package main

import "fmt"

func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}

func main() {
	arr := []int{254, 32, 64, 54, 39, 10, 86, 5, 319, 98}
	fmt.Println("Array before sorting : ", arr)
	selectionSort(arr)
	fmt.Println("Array after sorting: ", arr)
}
