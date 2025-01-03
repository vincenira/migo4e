// Create a function that concatenates two arrays into new array. Do not forget to test it with various types of input

package main

import "fmt"

func contenate_2_arrays_newArray(arr1 [4]int, arr2 [3]int) [7]int {
	contenated_slice := [7]int{}
	for i := 0; i < len(arr1); i++ {
		contenated_slice[i] = arr1[i]
	}

	for i := 0; i < len(arr2); i++ {
		contenated_slice[len(arr1)+i] = arr2[i]
	}
	return contenated_slice
}

func main() {
	array1 := [4]int{1, 3, 4, 6}
	array2 := [3]int{4, 91, 12}

	new_array := contenate_2_arrays_newArray(array1, array2)
	fmt.Println("capacity: ", cap(new_array), "length of the new slice: ", len(new_array))
	fmt.Println("new slice: ", new_array)
}
