package main

import "fmt"

func selection_sort(arr []int) {
	arr_len := len(arr)

	for i := 0; i < arr_len; i++ {
		lowest_index := i //Were tracking indicies and not numbers
		for j := i + 1; j < arr_len; j++ {
			if arr[j] < arr[lowest_index] {
				lowest_index = j // if the values at array[j] are smaller than val at arr[li] than the lowest index is now J
			}
		}
		if lowest_index != i {
			temp := arr[i]
			arr[i] = arr[lowest_index] // after this for loop if there were any vals that were lower than our current val were going to do a switcheroo
			arr[lowest_index] = temp
		}
	}
	fmt.Printf("Selection sorted list %v", arr)
}

func main() {
	unsorted_arr := []int{8, 4, 99, 234, 456, 274, 232, 687, 2}
	selection_sort(unsorted_arr)
}
