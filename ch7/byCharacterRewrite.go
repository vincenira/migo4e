/*
The bufio.Scanner in Go is designed to read input line by line, splitting it into tokens.
if you need to read a file character by character, a common aprroach is to use bufio.
NewReader in conjunction with Read() or ReadRune(). Implement the functionality of
byCharacter.go this way.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func charByChar(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		rChar, n, err := r.ReadRune()
		if err == io.EOF {
			if n != 0 {
				fmt.Println(string(rChar))
			}
			break
		} else if err != nil {
			fmt.Printf("Error reading file %s", err)
			return err
		}
		fmt.Println(string(rChar))

	}
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: byCharacter <file1> [<file2> ...]\n")
		return
	}

	for _, file := range args[1:] {
		err := charByChar(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
