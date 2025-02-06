/*
Go offers bufio.Scanner to read files line by line. Try to rewrite byLine.go using bufio.Scanner.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewScanner(f)
	for r.Scan() {
		fmt.Println(r.Text()) // Println will add back the final '\n'
	}
	if err := r.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "failed to read data from:", err)
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: byLine <file1> [<file2> ...]\n")
		return
	}

	for _, file := range args[1:] {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
