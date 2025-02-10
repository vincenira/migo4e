/*
Declaration of channels based on directions
1. Bidirectional channel : chan T
2. Send only channel: chan <- T
3. Receive only channel: <- chan T
*/
package main

import (
	"fmt"
)

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting")
	ch <- "Hello toto!"
	fmt.Println("Greeter completed")
}
func main() {
	msg := make(chan string)
	go greet(msg)
	greeting := <-msg
	fmt.Println("Greeting received!")
	fmt.Println(greeting)

}
