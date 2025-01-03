package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		// log.Fatal() function is used when something erroneous has happened and you just want to exit your program so soon as possible
		// after reporting that bad situation
		log.Fatal("Fatal: Hello World!")
	}
	// log.Panic() implies that something really unexpected and unknown, such as not being able to find a file that was previously accessed
	// or not having enough disk space, has happened.
	// it also includes additional low-level information that, hopefully, will help you resolve difficult situations that arise in your Go code.
	log.Panic("Panic: Hello World!")
}
