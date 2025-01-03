// Write a Go Program that converts an existing array into a map
package main

import (
	"fmt"
	"strconv"
)

func convertArrayToMap(array [5]int) map[string]int {
	m := make(map[string]int)
	for index, value := range array {
		m[strconv.Itoa(index)] = value
	}
	return m
}

func PrintMap(m map[string]int) {
	for key, value := range m {
		fmt.Printf("%s: %v\n", key, value)
	}
}

func main() {
	a1 := [5]int{1, 2, 3, 4, 5}
	a3 := [5]int{6, 7, 8, 9, 10}

	m1 := convertArrayToMap(a1)
	m3 := convertArrayToMap(a3)

	// Print the maps
	PrintMap(m1)
	PrintMap(m3)
}
