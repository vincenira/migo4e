/*
Write a Go utility that converts os.Args into a slice of structures, with fields for storing
the index and the value of each command line argument—you should define the structure
that is going to be used.
*/
package main

import (
	"fmt"
	"os"
)

type comandline struct {
	index int
	value string
}

func main() {
	args := os.Args
	length := len(args)
	if length < 2 {
		fmt.Println("please enter at least one argument")
		return
	}
	cmdlines := make([]comandline, 0, length)
	for i := range length - 1 {
		cmdlines = append(cmdlines, comandline{i + 1, args[i+1]})
	}

	for i := range length - 1 {
		fmt.Printf("index: %d, value: %s\n", cmdlines[i].index, cmdlines[i].value)
	}
}
