/*
Try to implement a concurrent version of wc(1) that uses shared memory.
We will implement it using the mutex to lock the variable to share.
*/
package main

import (
	"fmt"
	"os"
)

var read
func readFile(path string) {
}

func wordByWord(fileName string) {
}

func lineByLine(fileName string) {
}

func charByChar(fileName string) {
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: wcShaMem <file1> [<file2> ...]\n")
		return
	}

	for _, file := range args[1:] {
		err := wordByWord(file)
		if err != nil {
			fmt.Println(err)
		}

	}
}
