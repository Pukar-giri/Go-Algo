package main

import (
	"fmt"
)

func main() {
	//only for testing purposes
	arr := []int{5, 7, 9, 4, 1, 6, 8, 2, 3, 11, 23, 72, 52, 86, 8, 2, 81, 4, 64, 6, 16, 1, 31, 31, 31, 1, 31, 313, 131, 1, 54, 6, 646, 46, 64, 6, 66, 64, 4, 654, 654, 4, 4 + 6, 4, 46, +64, 64, 64, 64, 46, 46, 4 + 6, 46, 46, 465, 6, 64, 94, 65, 498, 161, 4, 9, 1, 4, 16, 65, 4, 6, 6, 64, 131, 31, 31, 16, 16, 11, 651, 651, 61, 1, 1, 16, 161, 61, 61, 1, 1, 1, 6, 6, 2, 65}
	fmt.Println(Bubblesort(arr))
}
func Bubblesort(arr []int) []int {
	//assume the array is not sorted
	sorted := false
	//until it is not sorted repeat doing the same
	for !sorted {
		//lets assume that we dont have to swap now as it is already sorted
		sorted = true
		//iterate over each of the elements and compare it with element next to it and swap if they are not in right order
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				tmp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
				//we had to swap this means array was not sorted
				sorted = false
			}
		}

	}
	return arr
}
