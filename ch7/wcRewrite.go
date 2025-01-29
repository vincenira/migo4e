/*
This code groupes 3 exercises of chapter 7.  it will be done in step with commit and comments to illustarte the changes
  - use the functionality of byCharacter.go, byLine.go, and byWord.go in order to create a simplified version of the wc(1)
    UNIX utility.
  - Create a full version of the wc(1) UNIX utility, using the viper package to process command line options
  - Create a full version of the wc(1) UNIX utility, using commmands instead of command line options, with the
    help of the cobra package
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
printResult takes two arguments
and then total only if the number of characters of the second arguments is zero
otherwise it prints both two arguments
*/
func countPerLinePerFile(s []string) (total int) {
	total = len(s)
	return
}

func countPerWordPerFile(s []string) (total int) {
	total = 0
	return
}

func countPerCharacterPerFile(s []string) (total int) {
	total = 0
	return
}

func printResult(total int, fileName string) {
	if len(fileName) == 0 {
		fmt.Println(total)
	} else {
		fmt.Printf("%d %s", total, fileName)
	}
}

func readfile(fileName string) ([]string, error) {
	var linesReaded []string
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		// ReadString() returns two values: the string that was read and an error variable.
		line, err := r.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				linesReaded = append(linesReaded, line)
			}
			break
		}

		if err != nil {
			fmt.Printf("error reading file %s", err)
			return nil, err
		}
		linesReaded = append(linesReaded, line)
	}
	return linesReaded, nil
}

func main() {
	args := os.Args
	lengthArgs := len(args)
	if lengthArgs > 1 {
		for index := range lengthArgs - 1 {
			s, _ := readfile(args[index+1])
			fmt.Println(s)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		var readlines []string
		for scanner.Scan() {
			readlines = append(readlines, scanner.Text())
		}
		fmt.Println(len(readlines))
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

	}
}
