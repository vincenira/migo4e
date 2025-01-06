/*
Go 1.21 comes with a new function named clear that clears maps and slices. For maps,
it deletes all entries whereas for slices it zeros all existing values. Experiment with it to
learn how it works
*/
package main

import "fmt"

func main() {
	sliceEx := []int{1, 2, 3, 4, 5}
	mapEx := make(map[string]int)

	mapEx = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println("before the usage of the clear function")
	fmt.Println(sliceEx)
	fmt.Println(mapEx)
	clear(sliceEx)
	clear(mapEx)
	fmt.Println("After the usage of the clear function")
	fmt.Println(sliceEx)
	fmt.Println(mapEx)
}
