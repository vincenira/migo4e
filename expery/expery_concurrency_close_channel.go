package main

import (
	"fmt"
	"time"
)

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting!")
	ch <- "Hello Toto"
	close(ch)
	fmt.Println("Greeter completed")
}

func main() {
	msg := make(chan string)
	go greet(msg)

	time.Sleep(2 * time.Second)
	greeting := <-msg

	time.Sleep(2 * time.Second)
	fmt.Println("Greeting received")
	fmt.Println(greeting)

	_, ok := <-msg
	if ok {
		fmt.Println("channel is open!")
	} else {
		fmt.Println("channel is closed!")
	}
}
