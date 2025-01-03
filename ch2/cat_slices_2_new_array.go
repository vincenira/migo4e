// create a function that concatenates two slices into a new array

package main

import "fmt"

func contenate_slices_2_newArray(sli1 []int, sli2 []int) [7]int {
	contenated_slice := [7]int{}
	for i := 0; i < len(sli1); i++ {
		contenated_slice[i] = sli1[i]
	}

	for i := 0; i < len(sli2); i++ {
		contenated_slice[len(sli1)+i] = sli2[i]
	}
	return contenated_slice
}

func main() {
	array1 := []int{1, 3, 4, 6}
	array2 := []int{4, 91, 12}

	new_array := contenate_slices_2_newArray(array1, array2)
	fmt.Println("capacity: ", cap(new_array), "length of the new array: ", len(new_array))
	fmt.Println("new array: ", new_array)
}
