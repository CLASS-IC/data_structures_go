package main

import "fmt"

func insertion_sort(arr []int) {
	arr_len := len(arr)
	for i := 1; i < arr_len; i++ {
		temp_val := arr[i]
		position := i
		for position > 0 { //while position is bigger than 0 meaning were going to iterate through previous positions
			if arr[position] < arr[position-1] { //if the num on right of array is bigger than left
				arr[position] = arr[position-1] //pull a switchero
				position = position - 1         //decrement position so that we can now compare our current val with anything to the left if possible
			}
			arr[position] = temp_val //set the next temp val as whats in the current position so we can compare
		}
		fmt.Printf("insertion sorted %v", arr)
	}

}

func main() {
	unsorted_array := []int{43, 2, 34, 23, 45, 62, 3, 6, 7, 29}
	insertion_sort(unsorted_array)
}
