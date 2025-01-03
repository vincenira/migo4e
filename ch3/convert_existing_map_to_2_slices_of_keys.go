/*
Write a Go program that converts an existing map into two slices—the first slice con-
taining the keys of the map whereas the second one containing the values. The values at
index n of the two slices should correspond to a key and value pair that can be found in
the original map.
*/
package main

import "fmt"

func convertMapTo2SlicesOfKeyValue(m map[string]int) ([]string, []int) {
	keys := make([]string, 0, len(m))
	values := make([]int, 0, len(m))
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func main() {
	M1 := map[string]int{"a": 1, "b": 3, "c": 4, "d": 5}
	M2 := map[string]int{"1": 12, "2": 10, "3": 13}

	keys1, values1 := convertMapTo2SlicesOfKeyValue(M1)
	keys2, values2 := convertMapTo2SlicesOfKeyValue(M2)

	fmt.Println(keys1)
	fmt.Println(values1)
	fmt.Println(keys2)
	fmt.Println(values2)
}
