// Run go doc errors Is in order to learn about errors.Is() and try to create a small Go
// program that uses it. After that, modify error.go to use errors.Is().
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Custom error message with errors.New()
var currentErr = errors.New("this is a custom error message")

func check(a, b int) error {
	if a == 0 && b == 0 {
		return currentErr
	}
	return nil
}

// Custom error message with fmt.Errorf()
func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Getuid())
	}
	return nil
}

func main() {
	err := check(0, 10)
	if err == nil {
		fmt.Println("check() executed normally!")
	} else {
		fmt.Println(err)
	}

	err = check(0, 0)
	// In order to be true the errors.Is has to point to the same error values.
	// if we used the errors.New() function the created err value will a different one even if the error message is the same.
	if errors.Is(err, currentErr) {
		fmt.Println("Custom error detected!")
	}

	err = formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("-123")
	if err == nil {
		fmt.Println("Int value is", i)
	}

	i, err = strconv.Atoi("Y123")
	if err != nil {
		fmt.Println(err)
	}
}
