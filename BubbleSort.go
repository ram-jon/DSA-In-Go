package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func main() {
	arr := []int{254, 32, 64, 54, 39, 10, 86, 5, 319, 98}
	fmt.Println("Array before sorting : ", arr)
	bubbleSort(arr)
	fmt.Println("Array after sorting: ", arr)
}
