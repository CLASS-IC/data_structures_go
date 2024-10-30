package main

import "fmt"

func bubble_sort(arr []int) {
	sorted := false
	lowest_index := len(arr) - 1
	for !sorted {
		sorted = true
		for i := 0; i < lowest_index; i++ {
			if arr[i] > arr[i+1] {
				sorted = false
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		lowest_index = lowest_index - 1
	}
	fmt.Printf("bubble sorted list %v", arr)
}

func main() {
	arr := []int{3, 6, 4, 87, 9, 2, 1, 46, 7}
	bubble_sort(arr)
}
