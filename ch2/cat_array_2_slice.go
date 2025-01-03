// Create and test a function that concatenates two arrays into a new slice.
package main

import "fmt"

func contenate_2_arrays(arr1 [4]int, arr2 [3]int) []int {
	contenated_slice := []int{}
	for i := 0; i < len(arr1); i++ {
		contenated_slice = append(contenated_slice, arr1[i])
	}

	for i := 0; i < len(arr2); i++ {
		contenated_slice = append(contenated_slice, arr2[i])
	}
	return contenated_slice
}

func main() {
	array1 := [4]int{1, 3, 4, 6}
	array2 := [3]int{4, 91, 12}

	new_slice := contenate_2_arrays(array1, array2)
	fmt.Println("capacity: ", cap(new_slice), "length of the new slice: ", len(new_slice))
	fmt.Println("new slice: ", new_slice)
}
