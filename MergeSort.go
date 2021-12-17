package main

import "fmt"

func merge(arr []int, si int, mid int, ei int) {
	larr := []int{}
	larr = append(larr, arr[si:mid+1]...)
	rarr := []int{}
	rarr = append(rarr, arr[mid+1:ei+1]...)
	oi := si
	li, ri := 0, 0
	for li < len(larr) && ri < len(rarr) {
		if larr[li] <= rarr[ri] {
			arr[oi] = larr[li]
			li++
		} else {
			arr[oi] = rarr[ri]
			ri++
		}
		oi++
	}
	for li < len(larr) {
		arr[oi] = larr[li]
		oi++
		li++
	}
	for ri < len(rarr) {
		arr[oi] = rarr[ri]
		oi++
		ri++
	}
}

func mergeSortHelper(arr []int, si int, ei int) {
	if si < ei {
		mid := int((si + ei) / 2)
		mergeSortHelper(arr, si, mid)
		mergeSortHelper(arr, mid+1, ei)
		merge(arr, si, mid, ei)
	}
}

func mergeSort(arr []int) {
	mergeSortHelper(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{254, 32, 64, 54, 39, 10, 86, 5, 319, 98}
	fmt.Println("Array before sorting : ", arr)
	mergeSort(arr)
	fmt.Println("Array after sorting: ", arr)
}
