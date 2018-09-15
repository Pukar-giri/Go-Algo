/*
*Author:-Pukar Giri
*Created On:-15Th September 2018 at 22:41
*File Name:-mergesort.go
*Project Name:-algorithms
*Licence:- MIT
*Email:-crazzy.lx75@gmail.com
 */

package main

import "fmt"

func main() {
	//just testing the algorithm here
	arr := []int{5, 7, 9, 4, 1, 6, 8, 2, 3, 11, 23, 72, 52, 86, 8, 2, 81, 4, 64, 6, 16, 1, 31, 31, 31, 1, 31, 313, 131, 1, 54, 6, 646, 46, 64, 6, 66, 64, 4, 654, 654, 4, 4 + 6, 4, 46, +64, 64, 64, 64, 46, 46, 4 + 6, 46, 46, 465, 6, 64, 94, 65, 498, 161, 4, 9, 1, 4, 16, 65, 4, 6, 6, 64, 131, 31, 31, 16, 16, 11, 651, 651, 61, 1, 1, 16, 161, 61, 61, 1, 1, 1, 6, 6, 2, 65}
	fmt.Println(arr)
	arr = MergeSort(arr)
	fmt.Println(arr)
}

//This is the best performing sorting algorithm for worse case scenarios
func MergeSort(arr []int) []int {
	//take the midpoint of the array
	mid := int(len(arr) / 2)
	//for all array of length more than 1 its sorted form is the  sorterd part before and after midpoint resorted smartly using merge function
	if len(arr) > 1 {
		return merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
	} else {
		//but a empty array and single element array are sorted by themselves
		return arr
	}
}

func merge(arr1 []int, arr2 []int) []int {
	//lets declare an empty array to store our result
	arr := []int{}
	// lets also count how many element from each array we have used so far
	used1 := 0
	used2 := 0
	//loop till either one of the array has been used completely
	for used1 < len(arr1) && used2 < len(arr2) {
		if arr1[used1] < arr2[used2] {
			//if next element to be used from first array is smaller than that of second the next element from
			// first will go into the result array
			//else next element will go from second array
			arr = append(arr, arr1[used1])
			used1++
		} else {
			arr = append(arr, arr2[used2])
			used2++

		}
	}
	//if second array is not used up completely then remaining element of second array are poured directly
	// as they are sorted among themselves and there is nothing to compare against.
	if used2 < len(arr2) {
		arr = append(arr, arr2[used2:]...)
	}
	//same is true for first array
	if used1 < len(arr1) {
		arr = append(arr, arr1[used1:]...)
	}
	//finally the result we made so far is returned
	return arr

}
