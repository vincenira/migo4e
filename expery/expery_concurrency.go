package main

import (
	"fmt"
	"time"
)

func helloworld() {
	fmt.Println("Hello World!")
}

func goodbye() {
	fmt.Println("Good Bye!")
}

func main() {
	go helloworld()
	//This time sleep function helps to allow enough time for the goroutine to
	// finish executing. if it is removed the goroutine won't result.
	// To experiment it removing the below time.Sleep function
	time.Sleep(1 * time.Second)
	goodbye()
}
