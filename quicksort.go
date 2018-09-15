/*
*Author:-Pukar Giri
*Created On:-15Th September 2018 at 19:55
*File Name:-quicksort.go
*Project Name:-algorithms
*Licence:- MIT
*Email:-crazzy.lx75@gmail.com
 */

package main

import "fmt"

func main() {
	//only for testing purposes
	arr := []int{5, 7, 9, 4, 1, 6, 8, 2, 3, 11, 23, 72, 52, 86, 8, 2, 81, 4, 64, 6, 16, 1, 31, 31, 31, 1, 31, 313, 131, 1, 54, 6, 646, 46, 64, 6, 66, 64, 4, 654, 654, 4, 4 + 6, 4, 46, +64, 64, 64, 64, 46, 46, 4 + 6, 46, 46, 465, 6, 64, 94, 65, 498, 161, 4, 9, 1, 4, 16, 65, 4, 6, 6, 64, 131, 31, 31, 16, 16, 11, 651, 651, 61, 1, 1, 16, 161, 61, 61, 1, 1, 1, 6, 6, 2, 65}
	fmt.Println(Quicksort(arr))
}

func Quicksort(arr []int) []int {
	//if the length of the array is 0 or 1 the array is sorted by itself
	if len(arr) <= 1 {
		return arr
	}
	//select a pivot element i chose to use the middle one
	pivot := arr[int(len(arr)/2)]
	//initialize a slice for elements that are less, equal or more than pivot
	eq := []int{}
	less := []int{}
	more := []int{}
	//iterate over the array and fill in the slices as declared above
	for i := 0; i < len(arr); i++ {
		if arr[i] == pivot {
			eq = append(eq, arr[i])
		} else if arr[i] < pivot {
			less = append(less, arr[i])
		} else if arr[i] > pivot {
			more = append(more, arr[i])
		}
	}
	//free the original array memory to save on memory
	arr = nil
	//fmt.Println(" less are",less,"equals are =",eq," more are",more)
	//the sorted array will be sorted less + equals +sorted more
	x := append(Quicksort(less), eq...)
	return append(x, Quicksort(more)...) //to arrange in descending order just swap the position of less and more in last two steps

}
