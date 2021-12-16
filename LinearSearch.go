package main

import "fmt"

func linearSearch(arr []int, x int) (bool, int) {
	for i, val := range arr {
		if val == x {
			return true, i
		}
	}
	return false, -1
}

func main() {
	arr := []int{10, 32, 54, 39, 86, 319, 98}
	x := 319
	if found, index := linearSearch(arr, x); found {
		fmt.Printf("Element %v found at index %v \n", x, index)
	} else {
		fmt.Println("Element not found!!")
	}
}
