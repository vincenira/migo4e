// ascii alphanumeric characters 33-126
// password generate

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var (
	MIN int = 0
	MAX int = 94
)

func random(m, M int) int {
	return rand.Intn(M-m) + m
}

func getString(stringlen int64) string {
	tempstr := ""
	refChar := "!"
	var i int64 = 0
	for i < stringlen {
		myRand := random(MIN, MAX)
		newchar := string(refChar[0] + byte(myRand))
		tempstr = tempstr + newchar
		i++
	}
	return tempstr
}

func main() {
	var LENGTH int64 = 8

	switch len(os.Args) {
	case 2:
		if n, err := strconv.ParseInt(os.Args[1], 10, 64); err == nil {
			if n > LENGTH {
				LENGTH = n
			}
		}
	default:
		fmt.Println("using the default")
	}
	fmt.Println(getString(LENGTH))
}
