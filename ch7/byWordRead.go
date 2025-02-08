/*
Similarly, byWord.go uses ReadString('\n') to read the input file. Modify the code to
use Scanner instead
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordByWord(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewScanner(f)
	r.Split(bufio.ScanWords)
	for r.Scan() {
		fmt.Println(r.Text()) // Println will add back the final '\n'
	}

	return nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: byWord <file1> [<file2> ...]\n")
		return
	}

	for _, file := range args[1:] {
		err := wordByWord(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
