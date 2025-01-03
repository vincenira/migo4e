/*
Change the code of intRE.go to process multiple command line arguments and display
totals of true and false results at the end.
*/

package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	numberOfTs := 0
	numberOfFs := 0
	for i := range len(arguments) - 1 {
		s := arguments[i+1]
		if matchInt(s) {
			numberOfTs += 1
		} else {
			numberOfFs += 1
		}
	}
	fmt.Printf("Total number of Trues: %d\nTotal number of Falses: %d\n", numberOfTs, numberOfFs)
}

/*
Testing code
go run intRE.go 1 3 4 -5 f t +5 f u

Total number of Trues: 5
Total number of Falses: 4
*/
