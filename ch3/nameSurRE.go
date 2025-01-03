/*
Make the necessary changes to nameSurRE.go to be able to process multiple command
line arguments.
*/

package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("at least one argument is required")
		return
	}

	for i := range len(arguments) - 1 {
		ret := matchNameSur(arguments[i+1])
		fmt.Println(ret)
	}
}

/*
testing code
go run nameSurRe.go test to po toto Test

false
false
false
false
true
*/
