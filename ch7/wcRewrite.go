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
	"os"
)

/*
printResult takes two arguments
and then total only if the number of characters of the second arguments is zero
otherwise it prints both two arguments
*/

func printResult(total int, fileName string) {
	if len(fileName) == 0 {
		fmt.Println(total)
	} else {
		fmt.Printf("%d %s", total, fileName)
	}
}

func main() {
	args := os.Args
	lengthArgs := len(args)
	if lengthArgs > 1 {
		for index := range lengthArgs - 1 {
			fmt.Println(args[index+1])
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text()) // Println will add back the final '\n'
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

	}
}
