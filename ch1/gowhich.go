package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
The gowhich is similar to the which linux command.
it is divided into three parts:
  - reading input argument, which is the name of the executable file
  - reading the value from the PATH environment variable, splitting it and iterating over these directories
  - looking for the desired binary fil
    if it can be found or not
  - regular file and  executable
  - return the ful path of the desired binary
*/
func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	// Get the name of the binary
	fileName := os.Args[1]

	// Get the current PATH value
	// should be merged the declaration with assignement on the next line but it is for learning purpose.
	var path string
	path = os.Getenv("PATH")

	var allOccurrences []string

	for _, directoryPath := range filepath.SplitList(path) {
		filePath := filepath.Join(directoryPath, fileName)

		fileDesc, err := os.Stat(filePath)
		if err != nil {
			continue
		}
		fileDescMode := fileDesc.Mode()

		if !fileDescMode.IsRegular() {
			continue
		}

		if fileDescMode&0111 != 0 {
			allOccurrences = append(allOccurrences, filePath)
		}
	}

	if len(allOccurrences) == 0 {
		fmt.Println("Not found")
	} else {
		for _, f := range allOccurrences {
			fmt.Println(f)
		}
		os.Exit(0)
	}
}
