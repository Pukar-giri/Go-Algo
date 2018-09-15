/*
*Author:-Pukar Giri
*Created On:-15Th September 2018 at 15:01
*File Name:-insertion sort.go
*Project Name:-algorithms
*Licence:- MIT
*Email:-crazzy.lx75@gmail.com
 */

package main

import "fmt"

func main() {
	//just testing the algorithm here
	arr := []int{5, 7, 9, 4, 1, 6, 8, 2, 3, 11, 23, 72, 52, 86, 8, 2, 81, 4, 64, 6, 16, 1, 31, 31, 31, 1, 31, 313, 131, 1, 54, 6, 646, 46, 64, 6, 66, 64, 4, 654, 654, 4, 4 + 6, 4, 46, +64, 64, 64, 64, 46, 46, 4 + 6, 46, 46, 465, 6, 64, 94, 65, 498, 161, 4, 9, 1, 4, 16, 65, 4, 6, 6, 64, 131, 31, 31, 16, 16, 11, 651, 651, 61, 1, 1, 16, 161, 61, 61, 1, 1, 1, 6, 6, 2, 65}
	arr = InsertionSort(arr)
	fmt.Println(arr)
}
func InsertionSort(arr []int) []int {
	//go from second to last element for each iteration lets call the element key
	for i := 1; i < len(arr); i++ {
		//select the key element and compare it against the element on left of it
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			//keep swaping until the element on left is smaller than the key element
			tmp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = tmp
			j--
		}
		//repeat the process for consecutive elements
	}
	return arr
}
