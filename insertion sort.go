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
	arr := []int{5, 7, 9, 4, 1, 6, 8, 2, 3, 11, 23, 72}
	arr = InsertionSort(arr)
	fmt.Println(arr)
}
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			tmp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = tmp
			j--
		}
	}
	return arr
}
