/*
The byLine.go utility uses ReadString('\n') to read the input file. Modify the code to
use Scanner (https://pkg.go.dev/bufio#Scanner) for reading.
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
